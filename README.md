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

## After adding part that somehow makes it all faster

| Test type        | Time     | Cycles |
|------------------|----------|--------|
| interfaced       |   1.831µs|   5824 |
| tabled           |   1.63µs |   5568 |
| direct tabled    |  1.695µs |   5930 |
| function pointed |  1.818µs |   6424 |


# UPDATE

When there are strings, but with strings.Builder:
```
   direct tabled, took:  1.091µs,     3828 cycles
function pointed, took:  1.174µs,     4126 cycles
          tabled, took:  1.075µs,     3769 cycles
      interfaced, took:  1.101µs,     3863 cycles
          direct, took:  1.069µs,     3750 cycles
```

```
goos: linux
goarch: amd64
pkg: cchp/logic
cpu: Intel(R) Core(TM) i3-8100 CPU @ 3.60GHz
BenchmarkDoInterfaced
BenchmarkDoInterfaced-4        	  976675	      1072 ns/op	     776 B/op	      25 allocs/op
BenchmarkDoTabled
BenchmarkDoTabled-4            	  974346	      1081 ns/op	     776 B/op	      25 allocs/op
BenchmarkDoDirectTabled
BenchmarkDoDirectTabled-4      	  990459	      1091 ns/op	     776 B/op	      25 allocs/op
BenchmarkDoFunctionPointed
BenchmarkDoFunctionPointed-4   	  887757	      1230 ns/op	     776 B/op	      25 allocs/op
BenchmarkDirect
BenchmarkDirect-4              	  939790	      1108 ns/op	     776 B/op	      25 allocs/op
PASS
ok  	cchp/logic	5.388s
```

When there are no strings:
```
   direct tabled, took:     54ns,       93 cycles
function pointed, took:     54ns,       91 cycles
          tabled, took:     54ns,       93 cycles
      interfaced, took:     53ns,       94 cycles
          direct, took:     33ns,       28 cycles
```

```
goos: linux
goarch: amd64
pkg: cchp/logic
cpu: Intel(R) Core(TM) i3-8100 CPU @ 3.60GHz
BenchmarkDoInterfaced
BenchmarkDoInterfaced-4        	55660905	        21.33 ns/op	       0 B/op	       0 allocs/op
BenchmarkDoTabled
BenchmarkDoTabled-4            	57802566	        21.01 ns/op	       0 B/op	       0 allocs/op
BenchmarkDoDirectTabled
BenchmarkDoDirectTabled-4      	58022000	        20.66 ns/op	       0 B/op	       0 allocs/op
BenchmarkDoFunctionPointed
BenchmarkDoFunctionPointed-4   	56561500	        21.06 ns/op	       0 B/op	       0 allocs/op
BenchmarkDirect
BenchmarkDirect-4              	476983138	         2.544 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	cchp/logic	6.351s
```

What could this mean is that if you are going to use dynamic dispatch / polymorphism for calculations
without using strings - don't do that. It's better to use direct calls to functions. But if you are work
with strings a lot, than it doesn't matter. Unless Casey will explain where I am wrong and how to fix it.