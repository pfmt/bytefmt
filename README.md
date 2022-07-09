# bytefmt

[![Build Status](https://cloud.drone.io/api/badges/pfmt/bytefmt/status.svg)](https://cloud.drone.io/pfmt/bytefmt)
[![Go Reference](https://pkg.go.dev/badge/github.com/pfmt/bytefmt.svg)](https://pkg.go.dev/github.com/pfmt/bytefmt)

fmt.Formatter for bytes for Go.  
Source files are distributed under the BSD-style license.

## About

The software is considered to be at a alpha level of readiness -
its extremely slow and allocates a lots of memory)

## Benchmark

```sh
$ go test -count=1 -race -bench ./... 
goos: linux
goarch: amd64
pkg: github.com/pfmt/bytefmt
cpu: 11th Gen Intel(R) Core(TM) i7-1165G7 @ 2.80GHz
BenchmarkNamesInitialize/bytefmt_test.go:25/B_K_M_G_T_P_E-8             39951830            27.57 ns/op
BenchmarkNamesUpdate/bytefmt_test.go:25/B_K_M_G_T_P_E-8                 26603611            45.11 ns/op
BenchmarkBytesKilobytes/bytefmt_test.go:135/more_than_one_kilobyte100500-8          140618920            8.509 ns/op
BenchmarkBytesMegabytes/bytefmt_test.go:192/more_than_one_megabyte10050000-8        138513890            8.809 ns/op
BenchmarkBytesGigabytes/bytefmt_test.go:249/more_than_one_gigabyte10050000000-8     142891880            8.600 ns/op
BenchmarkBytesTerabytes/bytefmt_test.go:305/more_than_one_terabyte10050000000000-8  134421912            8.848 ns/op
BenchmarkBytesPetabytes/bytefmt_test.go:361/more_than_one_petabyte10050000000000000-8           137171584            8.585 ns/op
BenchmarkBytesExabytes/bytefmt_test.go:417/more_than_one_exabyte10050000000000000000-8          115660819            9.995 ns/op
BenchmarkBytesString/bytefmt_test.go:478/more_than_one_kilobyte1128-8                             516264          2191 ns/op
BenchmarkBytesFormat/bytefmt_test.go:532/general_format_%_v_0-8                                   354811          3189 ns/op
BenchmarkBytesFormat/bytefmt_test.go:589/general_format_%_v_1-8                                   302640          3993 ns/op
BenchmarkBytesFormat/bytefmt_test.go:646/general_format_%_v_1024-8                                301698          3953 ns/op
BenchmarkBytesFormat/bytefmt_test.go:703/general_format_%_v_1128-8                                277060          3980 ns/op
BenchmarkBytesFormat/bytefmt_test.go:758/string_format_%_s_0-8                                    221612          5426 ns/op
BenchmarkBytesFormat/bytefmt_test.go:815/string_format_%_s_1-8                                    186770          6361 ns/op
BenchmarkBytesFormat/bytefmt_test.go:872/string_format_%_s_1024-8                                 175197          6228 ns/op
BenchmarkBytesFormat/bytefmt_test.go:929/string_format_%_s_1128-8                                 175212          6368 ns/op
BenchmarkBytesFormat/bytefmt_test.go:984/double-quoted_string_format_%_q_0-8                      182739          5915 ns/op
BenchmarkBytesFormat/bytefmt_test.go:1041/double-quoted_string_format_%_q_1-8                     174912          6732 ns/op
BenchmarkBytesFormat/bytefmt_test.go:1098/double-quoted_string_format_%_q_1024-8                  170857          6785 ns/op
BenchmarkBytesFormat/bytefmt_test.go:1155/double-quoted_string_format_%_q_1128-8                  151651          7537 ns/op
BenchmarkBytesFormat/bytefmt_test.go:1212/float_format_%_f_0-8                                    343336          3369 ns/op
BenchmarkBytesFormat/bytefmt_test.go:1261/float_format_%f_1-8                                     289450          4022 ns/op
BenchmarkBytesFormat/bytefmt_test.go:1316/float_format_%_f_1024-8                                 262078          4304 ns/op
BenchmarkBytesFormat/bytefmt_test.go:1371/float_format_%_f_1128-8                                 251532          4606 ns/op
BenchmarkBytesFormat/bytefmt_test.go:1426/float_precision_1_format_%_.1f_0-8                      321405          3671 ns/op
BenchmarkBytesFormat/bytefmt_test.go:1493/float_precision_1_format_%_.1f_1-8                      255786          4530 ns/op
BenchmarkBytesFormat/bytefmt_test.go:1548/float_precision_1_format_%_.1f_1024-8                   246115          4559 ns/op
BenchmarkBytesFormat/bytefmt_test.go:1603/float_precision_1_format_%_.1f_1128-8                   247501          4690 ns/op
BenchmarkBytesFormat/bytefmt_test.go:1658/integer_format_%_d_0-8                                  392388          2986 ns/op
BenchmarkBytesFormat/bytefmt_test.go:1713/integer_format_%_d_1-8                                  391381          2926 ns/op
BenchmarkBytesFormat/bytefmt_test.go:1768/integer_format_%_d_1024-8                               383541          2947 ns/op
BenchmarkBytesFormat/bytefmt_test.go:1823/integer_format_%_d_1128-8                               383028          2930 ns/op
BenchmarkBytesFormat/bytefmt_test.go:1878/integer_precision_2_format_%_.2d_0-8                    371394          3198 ns/op
BenchmarkBytesFormat/bytefmt_test.go:1951/integer_precision_2_format_%_.2d_1-8                    366048          3233 ns/op
BenchmarkBytesFormat/bytefmt_test.go:2030/integer_precision_2_format_%_.2d_1024-8                 370528          3233 ns/op
BenchmarkBytesFormat/bytefmt_test.go:2109/integer_precision_2_format_%_.2d_1128-8                 352330          3241 ns/op
BenchmarkBytesFormat/bytefmt_test.go:2189/name_%_d_1124-8                                         336912          3399 ns/op
PASS
ok      github.com/pfmt/bytefmt 53.015s
```
