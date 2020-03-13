BASE_PACKAGE_PATH=github.com/tecowl/slack-knocker
MAIN_PACKAGE_PATH=$(BASE_PACKAGE_PATH)/cmd/slack-knocker

.PHONY: build
build: $(GOX)
	$(GOX) \
		-osarch="darwin/amd64 linux/amd64 windows/386 windows/amd64" \
		-output="pkg/{{.Dir}}_{{.OS}}_{{.Arch}}" \
		$(MAIN_PACKAGE_PATH)

# https://github.com/mitchellh/gox
GOX=$(GOPATH)/bin/gox
$(GOX):
	go get github.com/mitchellh/gox

pkg:
	$(MAKE) build

VERSION=$(shell cat ./VERSION)

.PHONY: version
version:
	@echo $(VERSION)

# https://github.com/tcnksm/ghr
GHR=$(GOPATH)/bin/ghr
$(GHR):
	go get -u github.com/tcnksm/ghr

.PHONY: release
release: build $(GHR)
	@if [ "$(GITHUB_TOKEN)" = "" ]; then \
		echo "GITHUB_TOKEN is missing" && exit 1; \
	fi
	$(GHR) \
		v$(VERSION) \
		pkg
