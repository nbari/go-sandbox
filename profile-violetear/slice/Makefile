.PHONY: all bench

GO ?= go

all: bench

bench:
	${GO} test -run=^$$ -bench=.
