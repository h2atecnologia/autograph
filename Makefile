# This Source Code Form is subject to the terms of the Mozilla Public
# License, v. 2.0. If a copy of the MPL was not distributed with this
# file, You can obtain one at http://mozilla.org/MPL/2.0/.

PROJECT		:= github.com/mozilla-services/autograph
GO 			:= GO15VENDOREXPERIMENT=1 go
GOGETTER	:= GOPATH=$(shell pwd)/.tmpdeps go get -d
GOLINT 		:= golint

all: test lint vet generate install

install:
	$(GO) install github.com/mozilla-services/autograph

go_vendor_dependencies:
	$(GOGETTER) gopkg.in/yaml.v2
	$(GOGETTER) github.com/mozilla-services/go-mozlog
	$(GOGETTER) github.com/mozilla-services/hawk-go
	echo 'removing .git from vendored pkg and moving them to vendor'
	find .tmpdeps/src -type d -name ".git" ! -name ".gitignore" -exec rm -rf {} \; || exit 0
	cp -ar .tmpdeps/src/* vendor/
	rm -rf .tmpdeps

lint:
	$(GOLINT) $(PROJECT)

vet:
	$(GO) vet $(PROJECT)

test:
	$(GO) test github.com/mozilla-services/autograph

generate:
	$(GO) generate

.PHONY: all test generate clean autograph
