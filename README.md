# isnil

`nil` check that works for interfaces too. Example code and benchmarks.

Benchmarks:
```
#
# x != nil
#

# nil pointer
BenchmarkEqNilBasic
BenchmarkEqNilBasic-8           1000000000           0.247 ns/op

# nil interface
BenchmarkEqNilInterface
BenchmarkEqNilInterface-8       1000000000           0.247 ns/op


#
# IsNil(x)
#

# nil pointer
BenchmarkIsNilBasic
BenchmarkIsNilBasic-8           274689241            4.37 ns/op

# (*struct{})(nil)
BenchmarkIsNilInterface
BenchmarkIsNilInterface-8       277561260            4.26 ns/op

# nil
BenchmarkIsNilNil
BenchmarkIsNilNil-8             803470719            1.48 ns/op
```
