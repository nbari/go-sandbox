go tool pprof --alloc_space http://0:8080/debug/pprof/heap
go tool pprof --inuse_space http://0:8080/debug/pprof/heap

wrk -t4 -c100 -d10s http://0:8080/

torch:

    docker run uber/go-torch -u http://x.x.x.x:8080 -p > torch.svg
