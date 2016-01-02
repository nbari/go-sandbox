.PHONY: all get test build cover clean

GO ?= go

all: clean build

get:
	${GO} get

build: get
ifdef DEBUG
	${GO} get -u github.com/mailgun/godebug
	${GOPATH}/bin/godebug build
else
	${GO} build
endif

test: get
	${GO} test -v

clean:
	@rm -f overlap*

cover:
	${GO} test -cover && \
	${GO} test -coverprofile=coverage.out  && \
	${GO} tool cover -html=coverage.out
