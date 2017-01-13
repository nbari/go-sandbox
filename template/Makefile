.PHONY: all get test clean build generate

GO ?= go

all: clean build

generate:
	${GO} generate

build: generate
	${GO} build

clean:
	@rm -rf template
