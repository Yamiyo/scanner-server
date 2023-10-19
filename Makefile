GO_BUILD_FLAGS ?= -tags dynamic,static_all
GOLANGCI_LINT_BUILD_FLAGS ?= --build-tags musl
ifneq ($(filter arm%,$(shell uname -p)),)
	GO_BUILD_FLAGS = -tags musl,dynamic,static_all
	GOLANGCI_LINT_BUILD_FLAGS = --build-tags musl,dynamic
endif

GOROOT := $(shell go env GOROOT)
GOPATH := $(shell go env GOPATH)

CMD := $(wildcard ./cmd/*/)

OUTDIR := out
OUTBIN := $(OUTDIR)/bin
OUTRPT := $(OUTDIR)/report/unittest
OUTCVR := $(OUTDIR)/report/coverage

all: build
.PHONY: all

setup:
	go install github.com/jstemmer/go-junit-report@latest
	go install github.com/t-yuki/gocover-cobertura@latest
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(GOROOT)/bin
.PHONY: setup

build: $(foreach cmd,$(CMD),build@$(cmd))
.PHONY: build

define BUILD_CMD
build@$(1):
	mkdir -p $(OUTBIN)
	go build $(GO_BUILD_FLAGS) -o $(OUTBIN)/ $(1)
endef
.PHONY: build@$(1)
$(foreach cmd,$(CMD),$(eval $(call BUILD_CMD,$(cmd))))

clean:
	rm -fr $(OUTDIR)
.PHONY: clean

fumpt:
	gofumpt -l -w .
.PHONY: fumpt

lint:
	GOPATH=$(GOPATH) golangci-lint run $(GOLANGCI_LINT_BUILD_FLAGS) -v ./...
.PHONY: lint

test:
	go test $(GO_BUILD_FLAGS) ./...
.PHONY: test

junit-report: coverage
	cat $(OUTRPT)/app.log | go-junit-report > $(OUTRPT)/app.xml
	[ "`tail -n1 $(OUTRPT)/app.log`" != "FAIL" ] # stop CI if any test case failed
.PHONY: test-junit

coverage-html: coverage
	go tool cover -html="$(OUTCVR)/cover.out" -o "$(OUTCVR)/app.html"
.PHONY: coverage-html

coverage:
	mkdir -p $(OUTCVR) $(OUTRPT)
	go test -v $(GO_BUILD_FLAGS) -coverprofile $(OUTCVR)/cover.out -covermode count ./... 2>&1 | tee $(OUTRPT)/app.log
.PHONY: coverage
