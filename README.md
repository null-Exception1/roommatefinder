## Benchmark Results (Caching vs Non-Caching)

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



---
# next up
- worker pools implementation
- fan in/ fan out implementation
