# Go 성능 툴 테스트

## benchmark

```sh
    pprof

    ## web interface
    brew install graphviz
    go tool pprof -http=:8080 cpu.prof

    ## pprof text
    go tool pprof -top cpu.prof
    go tool pprof -text cpu.prof
```

- <a href="http://localhost:6060/debug/pprof/">web interface</a>
- <a href="http://localhost:6060/debug/pprof/profile"> cpu profile </a>
- <a href="http://localhost:6060/debug/pprof/heap"> heap profile </a>
- <a href="http://localhost:6060/debug/pprof/goroutine"> goroutine profile </a>

## buffer.go

### cpu

```sh
Type: cpu
Time: 2025-04-27 12:54:19 KST
Duration: 1.31s, Total samples = 770ms (58.94%)
Showing nodes accounting for 770ms, 100% of 770ms total
      flat  flat%   sum%        cum   cum%
     330ms 42.86% 42.86%      330ms 42.86%  syscall.syscall
     250ms 32.47% 75.32%      250ms 32.47%  runtime.pthread_cond_wait
     130ms 16.88% 92.21%      130ms 16.88%  runtime.pthread_cond_signal
      50ms  6.49% 98.70%       50ms  6.49%  runtime.usleep
      10ms  1.30%   100%       10ms  1.30%  runtime.semacquire1
```

### memory

```sh
Time: 2025-04-27 12:47:36 KST
Showing nodes accounting for 12292.32kB, 100% of 12292.32kB total
      flat  flat%   sum%        cum   cum%
 5120.23kB 41.65% 41.65%  5120.23kB 41.65%  github.com/google/uuid.UUID.String (inline)
    2052kB 16.69% 58.35%     2052kB 16.69%  runtime.allocm
 2048.03kB 16.66% 75.01%  2048.03kB 16.66%  github.com/google/uuid.NewRandomFromReader
 1024.02kB  8.33% 83.34%  1536.02kB 12.50%  github.com/zkfmapf123/go-buffer/src.JobQueue.Consumer
 1024.02kB  8.33% 91.67%  1024.02kB  8.33%  github.com/zkfmapf123/go-buffer/src.JobQueue.Producer
  512.01kB  4.17% 95.83%   512.01kB  4.17%  internal/sync.(*HashTrieMap[go.shape.struct { net/netip.isV6 bool; net/netip.zoneV6 string },go.shape.struct { weak._ [0]*go.shape.struct { net/netip.isV6 bool; net/netip.zoneV6 string }; weak.u unsafe.Pointer }]).All
  512.01kB  4.17%   100%   512.01kB  4.17%  github.com/zkfmapf123/go-buffer/src.process (inline)
```
