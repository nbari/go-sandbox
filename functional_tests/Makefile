.PHONY: all clean test build cover

GO ?= go

all: build test

build:
	${GO} build

clean:
	@rm -rf functional_tests *.out

test:
	${GO} test

cover:
	${GO} test -cover && \
	${GO} test -coverprofile=coverage.out  && \
	${GO} tool cover -html=coverage.out
