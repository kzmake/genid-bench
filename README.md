# genid benchmark
## Results
```bash
go test -benchmem -bench .
...

BenchmarkUUIDv1-12                       7264171               158.5 ns/op            48 B/op          1 allocs/op
BenchmarkUUIDv3-12                       3832998               275.8 ns/op           192 B/op          5 allocs/op
BenchmarkUUIDv4-12                       7610632               154.5 ns/op            64 B/op          2 allocs/op
BenchmarkUUIDv5-12                       3868934               312.3 ns/op           216 B/op          5 allocs/op
BenchmarkUUIDv6-12                       5467276               206.8 ns/op            56 B/op          2 allocs/op
BenchmarkUUIDv7-12                       5623354               213.2 ns/op            56 B/op          2 allocs/op
BenchmarkULID-12                         6552250               182.5 ns/op            16 B/op          1 allocs/op
BenchmarkXID-12                         11960138                94.23 ns/op            0 B/op          0 allocs/op
BenchmarkNanoID-12                       3999192               305.3 ns/op           144 B/op          3 allocs/op
BenchmarkKSUID-12                        3325003               367.6 ns/op             0 B/op          0 allocs/op
BenchmarkSandflake-12                    5518105               218.2 ns/op            35 B/op          2 allocs/op
BenchmarkSnowflake-12                    4923600               244.0 ns/op            24 B/op          1 allocs/op
BenchmarkSonyflake-12                      30091             39194 ns/op               0 B/op          0 allocs/op
BenchmarkShortUUID-12                     192013              5783 ns/op            2873 B/op        135 allocs/op
BenchmarkHashID-12                       1354794               900.0 ns/op           512 B/op          6 allocs/op
BenchmarkAutoIncrement-12                    283           4154550 ns/op            4171 B/op         53 allocs/op
BenchmarkAutoIncrementWithHashID-12          266           5877518 ns/op            4698 B/op         59 allocs/op
```
