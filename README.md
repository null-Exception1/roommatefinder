##
blog (for more details) - https://null-exception1.github.io/blog/posts/roommatefinder
## Stack
NextJS + Go + PostGresDB
## Benchmark results from different versions

### Environment
- **OS**: Linux  
- **Arch**: amd64  
- **CPU**: Intel(R) Core(TM) i7-9750H @ 2.60GHz  
- **Go pkg**: `golang/benchmarks/simplefetch`

---

NOTE: older benchmarks were removed for being a little flaky

## Worker Pool Benchmark Results

| NumWorkers | Requests/sec | ns/op   | B/op    | Allocs/op |
|------------|--------------|---------|---------|-----------|
| 1          | 2,705        | 369,634 | 13,353  | 124       |
| 50         | 25,360       | 39,433  | 10,325  | 85        |
| 100        | 13,000       | ~77,000 | (varies)| (varies)  |

### notes
- **throughput scales up** dramatically from 1 → 50 workers (almost 10×).  
- **diminishing returns** kick in after ~50 workers — at 100, throughput drops due to DB pool saturation, goroutine scheduling overhead, and contention.  
- **sweet spot** depends on CPU cores and DB connection pool size. On this machine (Intel i7‑9750H, 12 threads), ~50 workers gave peak throughput.


# Final production results of BlocksHandler


# No caching
```
Running tool: /usr/local/go/bin/go test -test.fullpath=true -benchmem -run=^$ -bench ^BenchmarkBlocksHandler$ golang/benchmarks/simplefetch

goos: linux
goarch: amd64
pkg: golang/benchmarks/simplefetch
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
=== RUN   BenchmarkBlocksHandler
BenchmarkBlocksHandler
CACHING:  false
CACHE HITS:  0
CACHE MISSES:  0
BenchmarkBlocksHandler-4             157           7382362 ns/op          293677 B/op       2777 allocs/op
PASS
ok      golang/benchmarks/simplefetch   1.166s
```


# With Caching
```
Running tool: /usr/local/go/bin/go test -test.fullpath=true -benchmem -run=^$ -bench ^BenchmarkBlocksHandler$ golang/benchmarks/simplefetch

goos: linux
goarch: amd64
pkg: golang/benchmarks/simplefetch
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
=== RUN   BenchmarkBlocksHandler
BenchmarkBlocksHandler
CACHING:  true
CACHE HITS:  87716
CACHE MISSES:  1
BenchmarkBlocksHandler-4            4177            286186 ns/op          197142 B/op       1681 allocs/op
PASS
ok      golang/benchmarks/simplefetch   1.204s
```

# Final production results of RoomsBlocksHandler

# No caching
```
Running tool: /usr/local/go/bin/go test -test.fullpath=true -benchmem -run=^$ -bench ^BenchmarkRoomsBlocksHandler$ golang/benchmarks/simplefetch

goos: linux
goarch: amd64
pkg: golang/benchmarks/simplefetch
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
=== RUN   BenchmarkRoomsBlocksHandler
BenchmarkRoomsBlocksHandler
BenchmarkRoomsBlocksHandler-4              10000            108929 ns/op              9180 req/s           14851 B/op             143 allocs/op
PASS
ok      golang/benchmarks/simplefetch   1.431s
```

# With caching

```
Running tool: /usr/local/go/bin/go test -test.fullpath=true -benchmem -run=^$ -bench ^BenchmarkRoomsBlocksHandler$ golang/benchmarks/simplefetch

goos: linux
goarch: amd64
pkg: golang/benchmarks/simplefetch
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
=== RUN   BenchmarkRoomsBlocksHandler
BenchmarkRoomsBlocksHandler
BenchmarkRoomsBlocksHandler-4              18542             87024 ns/op             11491 req/s           11587 B/op             101 allocs/op
PASS
ok      golang/benchmarks/simplefetch   2.968s
```
