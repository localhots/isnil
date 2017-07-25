## Checking for `nil` in Go

Benchmarks:
```
BenchmarkIsNilBasic-8           200000000            9.47 ns/op
BenchmarkIsNilInterface-8       200000000            9.18 ns/op
BenchmarkIsNilNil-8             1000000000           2.03 ns/op
```
