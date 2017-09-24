.PHONY: all clean profile test

all: clean test

test:
	 go test -run=^$$ -bench=.

profile:
	 go test -cpuprofile=cpu.out -benchmem -memprofile=mem.out -run=^$$ -bench=.

clean:
	rm -rf profile *.pdf *.prof *.test *.out
