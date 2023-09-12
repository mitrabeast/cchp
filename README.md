# Clean code, horrible performance in Golang

| Test type        | Time     | Cycles |
|------------------|----------|--------|
| interfaced       |   7.45µs |  26324 |
| tabled           |   1.63µs |   5568 |
| direct tabled    |  1.448µs |   4906 |
| function pointed |  1.538µs |   6188 |

## Benchmarks
```
goos: linux
goarch: amd64
pkg: cchp/logic
cpu: Intel(R) Core(TM) i3-8100 CPU @ 3.60GHz
BenchmarkDoInterface/interfaced-4         	1000000000	         0.0000030 ns/op
BenchmarkDoTabled/tabled-4                	1000000000	         0.0000028 ns/op
BenchmarkDoDirectTabled/direct_tabled-4   	1000000000	         0.0000038 ns/op
BenchmarkDoFunctionPointed/func_pointed-4 	1000000000	         0.0000040 ns/op
PASS
ok  	cchp/logic	0.008s
```