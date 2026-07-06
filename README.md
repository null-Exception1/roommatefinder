## Baseline Benchmarks

These results represent the raw performance of our Go handlers hitting Postgres directly, **without caching**.

### Environment
- **OS:** Linux (amd64)
- **CPU:** Intel(R) Core(TM) i7-9750H @ 2.60GHz
- **Go Version:** go1.22+
- **Database:** Postgres (Docker, seeded with ~1000 rows)

### Results (No cacheing)

| Benchmark                  | Iterations | Time per op (ns/op) | Memory (B/op) | Allocations/op |
|----------------------------|------------|---------------------|---------------|----------------|
| `BenchmarkBlocksHandler`   | 450        | 2,549,874           | 840,851       | 21,949         |
| `BenchmarkRoomsBlocksHandler` | 1,000,000,000 | 0.02023            | 0             | 0              |

