.PHONY: all view

all: clean build
#	wrk -c2 -d10s -t2 -s counter.lua http://0:8080
#	@awk 'NR <= 1 || !/^time_/' result.csv > result.csv.tmp && mv result.csv.tmp result.csv

view:
	jq -R -s -f csv2json.jq results/wrk.csv
	#  jq ".[].requestsPerSecond"

build:
	go build

clean:
	@rm -rf benchmark *.tmp *.csv results/*
