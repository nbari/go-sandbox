#!/bin/sh

START=$(date +%s)

write() {
    echo "\033[1;31m"
    echo \#----------------------------------------------------------------------------
    echo \# $1
    echo \#----------------------------------------------------------------------------'\033[0m'
}

write "100 sleeping 0ms"
./benchmark > results/go-stats.csv &
sleep 1
wrk -c100 -d15s -t2 -s counter.lua http://0:8080
pkill -9 benchmark
sleep 2

write "100 sleeping 10ms"
./benchmark -s 10 >> results/go-stats.csv &
sleep 1
wrk -c100 -d15s -t2 -s counter.lua http://0:8080
pkill -9 benchmark
sleep 2

write "100 sleeping 100ms"
./benchmark -s 100 >> results/go-stats.csv &
sleep 1
wrk -c100 -d15s -t2 -s counter.lua http://0:8080
pkill -9 benchmark
sleep 2

write "100 sleeping 500ms"
./benchmark -s 500 >> results/go-stats.csv &
sleep 1
wrk -c100 -d15s -t2 -s counter.lua http://0:8080
pkill -9 benchmark
sleep

write "300 sleeping 0ms"
./benchmark >> results/go-stats.csv &
sleep 1
wrk -c300 -d15s -t2 -s counter.lua http://0:8080
pkill -9 benchmark
sleep 2

write "300 sleeping 10ms"
./benchmark -s 10 >> results/go-stats.csv &
sleep 1
wrk -c300 -d15s -t2 -s counter.lua http://0:8080
pkill -9 benchmark
sleep 2

write "300 sleeping 100ms"
./benchmark -s 100 >> results/go-stats.csv &
sleep 1
wrk -c300 -d15s -t2 -s counter.lua http://0:8080
pkill -9 benchmark
sleep 2

write "300 sleeping 500ms"
./benchmark -s 500 >> results/go-stats.csv &
sleep 1
wrk -c300 -d15s -t2 -s counter.lua http://0:8080
pkill -9 benchmark
sleep 2

write "500 sleeping 0ms"
./benchmark >> results/go-stats.csv &
sleep 1
wrk -c500 -d15s -t2 -s counter.lua http://0:8080
pkill -9 benchmark
sleep 2

write "500 sleeping 10ms"
./benchmark -s 10 >> results/go-stats.csv &
sleep 1
wrk -c500 -d15s -t2 -s counter.lua http://0:8080
pkill -9 benchmark
sleep 2

write "500 sleeping 100ms"
./benchmark -s 100 >> results/go-stats.csv &
sleep 1
wrk -c500 -d15s -t2 -s counter.lua http://0:8080
pkill -9 benchmark
sleep 2

write "500 sleeping 500ms"
./benchmark -s 500 >> results/go-stats.csv &
sleep 1
wrk -c500 -d15s -t2 -s counter.lua http://0:8080
pkill -9 benchmark
sleep 2


END=$(date +%s)
DIFF=$(echo "$END - $START" | bc)
write "Done! in $DIFF seconds."

# clean
awk 'NR <= 1 || !/^date/' results/wrk.csv > result.csv.tmp && mv result.csv.tmp results/wrk.csv
awk 'NR <= 1 || !/^T/' results/go-stats.csv > go-stats.csv.tmp && mv go-stats.csv.tmp results/go-stats.csv
