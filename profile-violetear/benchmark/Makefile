.PHONY: all view

all: clean build
	@mkdir results

view:
	jq -R -s -f csv2json.jq results/wrk.csv

build:
	go build

clean:
	@rm -rf benchmark *.tmp *.txt *.csv results
