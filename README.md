##


## Benchmark results from different versions

### Environment
- **OS**: Linux  
- **Arch**: amd64  
- **CPU**: Intel(R) Core(TM) i7-9750H @ 2.60GHz  
- **Go pkg**: `golang/benchmarks/simplefetch`

---

### Without Caching
| Benchmark                  | Iterations | Time/op      | Bytes/op | Allocs/op |
|----------------------------|------------|--------------|----------|-----------|
| **BlocksHandler**          | 2877       | 388,047 ns   | 258,546  | 2207      |
| **RoomsBlocksHandler**     | 1,000,000,000 | 0.0009717 ns | 0        | 0         |

---

### With Caching
| Benchmark                  | Iterations | Time/op      | Bytes/op   | Allocs/op |
|----------------------------|------------|--------------|------------|-----------|
| **BlocksHandler**          | 28         | 54,649,916 ns | 17,629,019 | 459,798   |
| **RoomsBlocksHandler**     | 1,000,000,000 | 0.02232 ns  | 0          | 0         |

---

### Removed json.Marshal overhead (caching)

| Benchmark                  | Iterations | Time/op      | Bytes/op | Allocs/op |
|----------------------------|------------|--------------|----------|-----------|
| **BlocksHandler**          | 3780       | 287,733 ns   | 208,379  | 1723      |
| **RoomsBlocksHandler**     | 1,000,000,000 | 0.0003799 ns | 0        | 0         |

---

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


