# Clean code, horrible performance in Golang

| Test type        | Time     | Cycles |
|------------------|----------|--------|
| interfaced       |   7.45µs |  26324 |
| tabled           |   1.63µs |   5568 |
| direct tabled    |  1.448µs |   4906 |
| function pointed |  1.538µs |   6188 |