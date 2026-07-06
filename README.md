### Environment
- **OS**: Linux  
- **Arch**: amd64  
- **CPU**: Intel(R) Core(TM) i7-9750H @ 2.60GHz  
- **Go pkg**: `golang/benchmarks/simplefetch`

---

### Without Caching
| Benchmark                  | Iterations | Time/op      | Bytes/op | Allocs/op |
|----------------------------|------------|--------------|----------|-----------|
| **BlocksHandler**          | 3013       | 389,557 ns   | 258,563  | 2207      |
| **RoomsBlocksHandler**     | 1,000,000,000 | 0.001113 ns | 0        | 0         |

---

### With Caching
| Benchmark                  | Iterations | Time/op      | Bytes/op   | Allocs/op |
|----------------------------|------------|--------------|------------|-----------|
| **BlocksHandler**          | 30         | 40,127,427 ns | 17,628,630 | 459,797   |
| **RoomsBlocksHandler**     | 1,000,000,000 | 0.01648 ns  | 0          | 0         |

---

---

### Next Steps
- optimizing set updates by giving worker pools to golang using a fan in/fan out implementation maybe 
- really work to death on fan in/fan out implementation
