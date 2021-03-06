package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"

	"go.mozilla.org/autograph/formats"
	"go.mozilla.org/autograph/signer"
)

const monitorAuthID = "monitor"

// MonitoringInputData is the data signed by the monitoring handler
var MonitoringInputData = []byte(`AUTOGRAPH MONITORING`)

func (a *autographer) handleMonitor(w http.ResponseWriter, r *http.Request) {
	rid := getRequestID(r)
	starttime := time.Now()
	userid, err := a.authorize(r, []byte(""))
	if err != nil {
		httpError(w, r, http.StatusUnauthorized, "authorization verification failed: %v", err)
		return
	}
	if userid != monitorAuthID {
		httpError(w, r, http.StatusUnauthorized, "user is not permitted to call this endpoint")
		return
	}

	signers := a.getSigners()
	sigerrstrs := make([]string, len(signers))
	sigresps := make([]formats.SignatureResponse, len(signers))
	var wg sync.WaitGroup
	for i, s := range signers {
		wg.Add(1)

		go func(i int, s signer.Signer) {
			defer wg.Done()

			// First try the DataSigner interface. If the signer doesn't
			// implement it, try the FileSigner interface. If that's still
			// not implemented, return an error.
			if _, ok := s.(signer.DataSigner); ok {
				// sign with data set to the base64 of the string 'AUTOGRAPH MONITORING'
				sig, err := s.(signer.DataSigner).SignData(MonitoringInputData, s.(signer.DataSigner).GetDefaultOptions())
				if err != nil {
					sigerrstrs[i] = fmt.Sprintf("signing failed with error: %v", err)
					return
				}

				encodedsig, err := sig.Marshal()
				if err != nil {
					sigerrstrs[i] = fmt.Sprintf("encoding failed with error: %v", err)
					return
				}
				sigerrstrs[i] = ""
				sigresps[i] = formats.SignatureResponse{
					Ref:        id(),
					Type:       s.Config().Type,
					Mode:       s.Config().Mode,
					SignerID:   s.Config().ID,
					PublicKey:  s.Config().PublicKey,
					Signature:  encodedsig,
					X5U:        s.Config().X5U,
					SignerOpts: s.Config().SignerOpts,
				}
				return
			}

			if _, ok := s.(signer.FileSigner); ok {
				// Signers that only implement the FileSigner interface must
				// also implement the TestFileGetter interface to return a valid
				// test file that can be used here to monitor the signer.
				if _, ok := s.(signer.TestFileGetter); !ok {
					sigerrstrs[i] = fmt.Sprintf("signer %q implements FileSigner but not the TestFileGetter interface", s.Config().ID)
					return
				}
				output, err := s.(signer.FileSigner).SignFile(s.(signer.TestFileGetter).GetTestFile(), s.(signer.FileSigner).GetDefaultOptions())
				if err != nil {
					sigerrstrs[i] = fmt.Sprintf("signing failed with error: %v", err)
					return
				}
				signedfile := base64.StdEncoding.EncodeToString(output)
				sigerrstrs[i] = ""
				sigresps[i] = formats.SignatureResponse{
					Ref:        id(),
					Type:       s.Config().Type,
					Mode:       s.Config().Mode,
					SignerID:   s.Config().ID,
					PublicKey:  s.Config().PublicKey,
					SignedFile: signedfile,
					X5U:        s.Config().X5U,
					SignerOpts: s.Config().SignerOpts,
				}
				return
			}

			sigerrstrs[i] = fmt.Sprintf("signer %q does not implement DataSigner or FileSigner interfaces", s.Config().ID)
			return
		}(i, s)
	}
	wg.Wait()
	for _, errstr := range sigerrstrs {
		if errstr != "" {
			httpError(w, r, http.StatusInternalServerError, errstr)
			return
		}
	}

	respdata, err := json.Marshal(sigresps)
	if err != nil {
		httpError(w, r, http.StatusInternalServerError, "signing failed with error: %v", err)
		return
	}
	if a.debug {
		log.Printf("signature response: %s", respdata)
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(respdata)
	log.WithFields(log.Fields{
		"rid":     rid,
		"user_id": userid,
		"t":       int32(time.Since(starttime) / time.Millisecond), //  request processing time in ms
	}).Info("monitoring operation succeeded")
}
