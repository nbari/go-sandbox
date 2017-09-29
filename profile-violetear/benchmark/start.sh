#!/bin/sh

START=$(date +%s)

write() {
    echo "\033[1;31m"
    echo \#----------------------------------------------------------------------------
    echo \# $1
    echo \#----------------------------------------------------------------------------'\033[0m'
}

# 0 ms
write "sleeping 0ms"
./benchmark &
sleep 1
wrk -c2 -d15s -t2 -s counter.lua http://0:8080
pkill -9 benchmark
sleep 2

write "sleeping 10ms"
./benchmark -s 10 &
sleep 1
wrk -c2 -d15s -t2 -s counter.lua http://0:8080
pkill -9 benchmark
sleep 2

write "sleeping 100ms"
./benchmark -s 100 &
sleep 1
wrk -c2 -d15s -t2 -s counter.lua http://0:8080
pkill -9 benchmark
sleep 2

write "sleeping 500ms"
./benchmark -s 500 &
sleep 1
wrk -c2 -d15s -t2 -s counter.lua http://0:8080
pkill -9 benchmark
sleep 2

END=$(date +%s)
DIFF=$(echo "$END - $START" | bc)
write "Done! build in $DIFF seconds."

# clean
awk 'NR <= 1 || !/^time_/' results/wrk.csv > result.csv.tmp && mv result.csv.tmp results/wrk.csv
awk 'NR <= 1 || !/^T/' results/go-stats.csv > go-stats.csv.tmp && mv go-stats.csv.tmp results/go-stats.csv
