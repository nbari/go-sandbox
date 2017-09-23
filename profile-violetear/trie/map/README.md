go test -cpuprofile cpu.prof -bench=.

go test -memprofile mem.prof -bench=.

go tool pprof cpu.prof
