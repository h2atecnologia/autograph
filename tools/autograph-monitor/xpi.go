package main

import (
	"log"

	"go.mozilla.org/autograph/signer/xpi"
)

func verifyXPISignature(sig string) error {
	xpiSig, err := xpi.Unmarshal(sig, []byte(inputdata))
	if err != nil {
		log.Fatal(err)
	}
	err = xpiSig.VerifyWithChain(conf.truststore)
	if err == nil {
		return nil
	}
	log.Printf("Got error %s verifying XPI signature with rel truststore trying dep truststore", err)
	return xpiSig.VerifyWithChain(conf.depTruststore)
}

const firefoxPkiStageRootHash = `DB:74:CE:58:E4:F9:D0:9E:E0:42:36:BE:6C:C5:C4:F6:6A:E7:74:7D:C0:21:42:7A:03:BC:2F:57:0C:8B:9B:90`

const firefoxPkiStageRoot = `-----BEGIN CERTIFICATE-----
MIIHYzCCBUugAwIBAgIBATANBgkqhkiG9w0BAQwFADCBqDELMAkGA1UEBhMCVVMx
CzAJBgNVBAgTAkNBMRYwFAYDVQQHEw1Nb3VudGFpbiBWaWV3MRwwGgYDVQQKExNB
ZGRvbnMgVGVzdCBTaWduaW5nMSQwIgYDVQQDExt0ZXN0LmFkZG9ucy5zaWduaW5n
LnJvb3QuY2ExMDAuBgkqhkiG9w0BCQEWIW9wc2VjK3N0YWdlcm9vdGFkZG9uc0Bt
b3ppbGxhLmNvbTAeFw0xNTAyMTAxNTI4NTFaFw0yNTAyMDcxNTI4NTFaMIGoMQsw
CQYDVQQGEwJVUzELMAkGA1UECBMCQ0ExFjAUBgNVBAcTDU1vdW50YWluIFZpZXcx
HDAaBgNVBAoTE0FkZG9ucyBUZXN0IFNpZ25pbmcxJDAiBgNVBAMTG3Rlc3QuYWRk
b25zLnNpZ25pbmcucm9vdC5jYTEwMC4GCSqGSIb3DQEJARYhb3BzZWMrc3RhZ2Vy
b290YWRkb25zQG1vemlsbGEuY29tMIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIIC
CgKCAgEAv/OSHh5uUMMKKuBh83kikuJ+BW4fQCHVZvADZh2qHNH8pSaME/YqMItP
5XQ1N5oLq1tRQO77AKn+eYPDAQkg+9VV+ct4u76YctcU/gvjieGKQ0fvuDH18QLD
hqa4DHgDmpCa/w+Eqzd54HaFj7ew9Bb7GZPHuZfk7Ct9fcN6kHneEj3KeuLiqzSV
VCRFV9RTlrUdsc1/VwF4A97JTXc3HJeWJO3azOlFpaJ8QHhmgXLLmB59HPeZ10Sf
9QwVGaKcn7yLuwtIA+wDhs8iwGZWcgmknW4DkkRDbQo7L+//4kVK+Yqq0HamZArm
vE4xENvbwOze4XYkCO3PwgmCotU7K5D3sMUUxkOaodlemO9OqRW8vJOJH3b6mhST
aunQR9/GOJ7sl4egrn2fOVZhBvM29lyBCKBffeQgtIMcKpeEKa4TNx4nTrWu1J9k
jHlvNeVL3FzMzJXRPl0RV71cYak+G6GnQ4fg3+4ZSSPxTvbwRJAO2xajkURxFSZo
sXcjYG8iPTSrDazj4LN2+882t4Q2/rMYpkowwLGbvJqHiw2tg9/hpLn1K4W18vcC
vFgzNRrTdKaJ/KjD17eJl8s8oPA7TiophPeezy1WzAc4mdlXS6A85b0mKDDU2A/4
3YmltjsSmizR2LnfeNs125EsCWxSUrAsnUYRO+lJOyNr7GGKGscCAwZVN6OCAZQw
ggGQMAwGA1UdEwQFMAMBAf8wDgYDVR0PAQH/BAQDAgEGMBYGA1UdJQEB/wQMMAoG
CCsGAQUFBwMDMCwGCWCGSAGG+EIBDQQfFh1PcGVuU1NMIEdlbmVyYXRlZCBDZXJ0
aWZpY2F0ZTAzBglghkgBhvhCAQQEJhYkaHR0cDovL2FkZG9ucy5tb3ppbGxhLm9y
Zy9jYS9jcmwucGVtMB0GA1UdDgQWBBSE6l/Nb0ySL+rR9PXIo7LCDLqm9jCB1QYD
VR0jBIHNMIHKgBSE6l/Nb0ySL+rR9PXIo7LCDLqm9qGBrqSBqzCBqDELMAkGA1UE
BhMCVVMxCzAJBgNVBAgTAkNBMRYwFAYDVQQHEw1Nb3VudGFpbiBWaWV3MRwwGgYD
VQQKExNBZGRvbnMgVGVzdCBTaWduaW5nMSQwIgYDVQQDExt0ZXN0LmFkZG9ucy5z
aWduaW5nLnJvb3QuY2ExMDAuBgkqhkiG9w0BCQEWIW9wc2VjK3N0YWdlcm9vdGFk
ZG9uc0Btb3ppbGxhLmNvbYIBATANBgkqhkiG9w0BAQwFAAOCAgEAck21RaAcTzbT
vmqqcCezBd5Gej6jV53HItXfF06tLLzAxKIU1loLH/330xDdOGyiJdvUATDVn8q6
5v4Kae2awON6ytWZp9b0sRdtlLsRo8EWOoRszCqiMWdl1gnGMaV7e2ycz/tR+PoK
GxHCh8rbOtG0eiVJIyRijLDjtExW8Eg+uz6Zkg1IWXqInj7Gqr23FOqD76uAfE82
YTWW3lzxpP3gL7pmV5G7ob/tIyAfrPEB4w0Nt2HEl9h7NDtKPMprrOLPkrI9eAVU
QeeI3RpAKnXOFQkqPYPXIlAaJ6qxtYa6tWHOqRyS1xKnvy/uWjEtU3tYJ5eUL1+2
vzNTdakJgkZDRdDNg0V3NYwza6BwL80VPSfqc1H6R8CU1uj+kjTlCEsoTPLeW7k5
t+lKHFMj0HZLNymgDD5f9UpI7yiOAIF0z4WKAMv/f12vnAPwmOPuOikRNOv0nNuL
RIpKO53Cd7aV5PdB0pNSPNjc6V+5IPrepALNQhKIpzoHA4oG+LlVVy4R3csPcj4e
zQQ9gt3NC2OXF4hveHfKZdCnb+BBl4S71QMYYCCTe+EDCsIGuyXWD/K2hfLD8TPW
thPX5WNsS8bwno2ccqncVLQ4PZxOIB83DFBFmAvTuBiAYWq874rneTXqInHyeCq+
819l9s72pDsFaGevmm0Us9bYuufTS5U=
-----END CERTIFICATE-----`

const firefoxPkiProdRootHash = `97:E8:BA:9C:F1:2F:B3:DE:53:CC:42:A4:E6:57:7E:D6:4D:F4:93:C2:47:B4:14:FE:A0:36:81:8D:38:23:56:0E`
const firefoxPkiProdRoot = `-----BEGIN CERTIFICATE-----
MIIGYTCCBEmgAwIBAgIBATANBgkqhkiG9w0BAQwFADB9MQswCQYDVQQGEwJVUzEc
MBoGA1UEChMTTW96aWxsYSBDb3Jwb3JhdGlvbjEvMC0GA1UECxMmTW96aWxsYSBB
TU8gUHJvZHVjdGlvbiBTaWduaW5nIFNlcnZpY2UxHzAdBgNVBAMTFnJvb3QtY2Et
cHJvZHVjdGlvbi1hbW8wHhcNMTUwMzE3MjI1MzU3WhcNMjUwMzE0MjI1MzU3WjB9
MQswCQYDVQQGEwJVUzEcMBoGA1UEChMTTW96aWxsYSBDb3Jwb3JhdGlvbjEvMC0G
A1UECxMmTW96aWxsYSBBTU8gUHJvZHVjdGlvbiBTaWduaW5nIFNlcnZpY2UxHzAd
BgNVBAMTFnJvb3QtY2EtcHJvZHVjdGlvbi1hbW8wggIgMA0GCSqGSIb3DQEBAQUA
A4ICDQAwggIIAoICAQC0u2HXXbrwy36+MPeKf5jgoASMfMNz7mJWBecJgvlTf4hH
JbLzMPsIUauzI9GEpLfHdZ6wzSyFOb4AM+D1mxAWhuZJ3MDAJOf3B1Rs6QorHrl8
qqlNtPGqepnpNJcLo7JsSqqE3NUm72MgqIHRgTRsqUs+7LIPGe7262U+N/T0LPYV
Le4rZ2RDHoaZhYY7a9+49mHOI/g2YFB+9yZjE+XdplT2kBgA4P8db7i7I0tIi4b0
B0N6y9MhL+CRZJyxdFe2wBykJX14LsheKsM1azHjZO56SKNrW8VAJTLkpRxCmsiT
r08fnPyDKmaeZ0BtsugicdipcZpXriIGmsZbI12q5yuwjSELdkDV6Uajo2n+2ws5
uXrP342X71WiWhC/dF5dz1LKtjBdmUkxaQMOP/uhtXEKBrZo1ounDRQx1j7+SkQ4
BEwjB3SEtr7XDWGOcOIkoJZWPACfBLC3PJCBWjTAyBlud0C5n3Cy9regAAnOIqI1
t16GU2laRh7elJ7gPRNgQgwLXeZcFxw6wvyiEcmCjOEQ6PM8UQjthOsKlszMhlKw
vjyOGDoztkqSBy/v+Asx7OW2Q7rlVfKarL0mREZdSMfoy3zTgtMVCM0vhNl6zcvf
5HNNopoEdg5yuXo2chZ1p1J+q86b0G5yJRMeT2+iOVY2EQ37tHrqUURncCy4uwIB
A6OB7TCB6jAMBgNVHRMEBTADAQH/MA4GA1UdDwEB/wQEAwIBBjAWBgNVHSUBAf8E
DDAKBggrBgEFBQcDAzCBkgYDVR0jBIGKMIGHoYGBpH8wfTELMAkGA1UEBhMCVVMx
HDAaBgNVBAoTE01vemlsbGEgQ29ycG9yYXRpb24xLzAtBgNVBAsTJk1vemlsbGEg
QU1PIFByb2R1Y3Rpb24gU2lnbmluZyBTZXJ2aWNlMR8wHQYDVQQDExZyb290LWNh
LXByb2R1Y3Rpb24tYW1vggEBMB0GA1UdDgQWBBSzvOpYdKvhbngqsqucIx6oYyyX
tzANBgkqhkiG9w0BAQwFAAOCAgEAaNSRYAaECAePQFyfk12kl8UPLh8hBNidP2H6
KT6O0vCVBjxmMrwr8Aqz6NL+TgdPmGRPDDLPDpDJTdWzdj7khAjxqWYhutACTew5
eWEaAzyErbKQl+duKvtThhV2p6F6YHJ2vutu4KIciOMKB8dslIqIQr90IX2Usljq
8Ttdyf+GhUmazqLtoB0GOuESEqT4unX6X7vSGu1oLV20t7t5eCnMMYD67ZBn0YIU
/cm/+pan66hHrja+NeDGF8wabJxdqKItCS3p3GN1zUGuJKrLykxqbOp/21byAGog
Z1amhz6NHUcfE6jki7sM7LHjPostU5ZWs3PEfVVgha9fZUhOrIDsyXEpCWVa3481
LlAq3GiUMKZ5DVRh9/Nvm4NwrTfB3QkQQJCwfXvO9pwnPKtISYkZUqhEqvXk5nBg
QCkDSLDjXTx39naBBGIVIqBtKKuVTla9enngdq692xX/CgO6QJVrwpqdGjebj5P8
5fNZPABzTezG3Uls5Vp+4iIWVAEDkK23cUj3c/HhE+Oo7kxfUeu5Y1ZV3qr61+6t
ZARKjbu1TuYQHf0fs+GwID8zeLc2zJL7UzcHFwwQ6Nda9OJN4uPAuC/BKaIpxCLL
26b24/tRam4SJjqpiq20lynhUrmTtt6hbG3E1Hpy3bmkt2DYnuMFwEx2gfXNcnbT
wNuvFqc=
-----END CERTIFICATE-----`
