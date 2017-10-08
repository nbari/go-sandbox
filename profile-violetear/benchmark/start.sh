#!/bin/sh

START=$(date +%s)

write() {
    echo "\033[1;31m"
    echo \#----------------------------------------------------------------------------
    echo \# $1
    echo \#----------------------------------------------------------------------------'\033[0m'
}

for t in 0 10 100 500
do
write "100 sleeping ${t}ms"
./benchmark -s ${t} >> results/go-stats.csv &
sleep 1
wrk -c100 -d15s -t2 -s counter.lua http://0:8080
pkill -9 benchmark
sleep 2
done

for t in 0 10 100 500
do
write "300 sleeping ${t}ms"
./benchmark -s ${t} >> results/go-stats.csv &
sleep 1
wrk -c300 -d15s -t2 -s counter.lua http://0:8080
pkill -9 benchmark
sleep 2
done

for t in 0 10 100 500
do
write "500 sleeping ${t}ms"
./benchmark -s ${t} >> results/go-stats.csv &
sleep 1
wrk -c500 -d15s -t2 -s counter.lua http://0:8080
pkill -9 benchmark
sleep 2
done

END=$(date +%s)
DIFF=$(echo "$END - $START" | bc)
write "Done! in $DIFF seconds."

# clean
awk 'NR <= 1 || !/^date/' results/wrk.csv > result.csv.tmp && mv result.csv.tmp results/wrk.csv
awk 'NR <= 1 || !/^T/' results/go-stats.csv > go-stats.csv.tmp && mv go-stats.csv.tmp results/go-stats.csv
