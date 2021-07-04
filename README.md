# genid benchmark
## Results
```bash
go test -benchmem -bench .
...

BenchmarkUUIDv1NewString-12                              7106696               158.5 ns/op            48 B/op          1 allocs/op
BenchmarkUUIDv3NewString-12                              3540658               284.0 ns/op           192 B/op          5 allocs/op
BenchmarkUUIDv4NewString-12                              7283752               158.7 ns/op            64 B/op          2 allocs/op
BenchmarkUUIDv5NewString-12                              3275497               351.1 ns/op           216 B/op          5 allocs/op
BenchmarkUUIDv6NewString-12                              5669246               226.3 ns/op            56 B/op          2 allocs/op
BenchmarkUUIDv7NewString-12                              5047783               205.5 ns/op            56 B/op          2 allocs/op
BenchmarkSnowFlakeNewUint64-12                             30070             39285 ns/op               0 B/op          0 allocs/op
BenchmarkULIDNewString-12                                5410569               191.0 ns/op            16 B/op          1 allocs/op
BenchmarkXIDNewString-12                                11135995               106.6 ns/op             0 B/op          0 allocs/op
BenchmarkShortUUIDNewString-12                            205339              5808 ns/op            2873 B/op        135 allocs/op
BenchmarkHashIDEncodeString-12                           1233394               952.1 ns/op           512 B/op          6 allocs/op
BenchmarkAutoIncrementString-12                              314           3715593 ns/op            4182 B/op         53 allocs/op
BenchmarkAutoIncrementWithHashIDEncodeString-12              364           3561762 ns/op            4691 B/op         59 allocs/op
```
