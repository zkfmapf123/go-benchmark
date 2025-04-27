run:
	go run main.go

bm-concurrency:
	go test -bench=BenchmarkConcurrent ./src

################################ Profile ################################
bm-cpu:
	go test -bench=. -cpuprofile=./benchmark/cpu.prof ./src

bm-mem:
	go test -bench=. -memprofile=./benchmark/mem.prof ./src

bm-good-cpu:
	go test -bench=. -cpuprofile=./benchmark/good-cpu.prof ./src

bm-good-mem:
	go test -bench=. -memprofile=./benchmark/good-mem.prof ./src

pp-cpu:
	go tool pprof cpu.prof

pp-mem:
	go tool pprof mem.prof
