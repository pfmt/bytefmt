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
BenchmarkNamesInitialize/bytefmt_test.go:25/B_K_M_G_T_P_E-8             41843841            26.96 ns/op
BenchmarkNamesUpdate/bytefmt_test.go:25/B_K_M_G_T_P_E-8                 26456245            44.91 ns/op
BenchmarkBytesKilobytes/bytefmt_test.go:135/more_than_one_kilobyte100500-8          137715493            8.765 ns/op
BenchmarkBytesMegabytes/bytefmt_test.go:192/more_than_one_megabyte10050000-8        126833665            9.423 ns/op
BenchmarkBytesGigabytes/bytefmt_test.go:249/more_than_one_gigabyte10050000000-8     135554748            9.539 ns/op
BenchmarkBytesTerabytes/bytefmt_test.go:305/more_than_one_terabyte10050000000000-8  124303792            9.680 ns/op
BenchmarkBytesPetabytes/bytefmt_test.go:361/more_than_one_petabyte10050000000000000-8           121375033            9.798 ns/op
BenchmarkBytesExabytes/bytefmt_test.go:417/more_than_one_exabyte10050000000000000000-8          125548881            8.874 ns/op
BenchmarkBytesString/bytefmt_test.go:478/more_than_one_kilobyte1128-8                             504247          2123 ns/op
BenchmarkBytesFormat/bytefmt_test.go:532/general_format_%_v_0-8                                   364606          3199 ns/op
BenchmarkBytesFormat/bytefmt_test.go:589/general_format_%_v_1-8                                   283976          4016 ns/op
BenchmarkBytesFormat/bytefmt_test.go:646/general_format_%_v_1024-8                                284350          4023 ns/op
BenchmarkBytesFormat/bytefmt_test.go:703/general_format_%_v_1128-8                                284601          4017 ns/op
BenchmarkBytesFormat/bytefmt_test.go:758/string_format_%_s_0-8                                    200888          5449 ns/op
BenchmarkBytesFormat/bytefmt_test.go:815/string_format_%_s_1-8                                    188862          6303 ns/op
BenchmarkBytesFormat/bytefmt_test.go:872/string_format_%_s_1024-8                                 187564          6231 ns/op
BenchmarkBytesFormat/bytefmt_test.go:929/string_format_%_s_1128-8                                 174298          6419 ns/op
BenchmarkBytesFormat/bytefmt_test.go:984/double-quoted_string_format_%_q_0-8                      182583          6004 ns/op
BenchmarkBytesFormat/bytefmt_test.go:1041/double-quoted_string_format_%_q_1-8                     161416          6757 ns/op
BenchmarkBytesFormat/bytefmt_test.go:1098/double-quoted_string_format_%_q_1024-8                  165452          6782 ns/op
BenchmarkBytesFormat/bytefmt_test.go:1155/double-quoted_string_format_%_q_1128-8                  142417          7734 ns/op
BenchmarkBytesFormat/bytefmt_test.go:1212/float_format_%_f_0-8                                    345793          3461 ns/op
BenchmarkBytesFormat/bytefmt_test.go:1261/float_format_%f_1-8                                     288199          4000 ns/op
BenchmarkBytesFormat/bytefmt_test.go:1316/float_format_%_f_1024-8                                 269931          4403 ns/op
BenchmarkBytesFormat/bytefmt_test.go:1371/float_format_%_f_1128-8                                 246644          4593 ns/op
BenchmarkBytesFormat/bytefmt_test.go:1426/float_precision_1_format_%_.1f_0-8                      320864          3626 ns/op
BenchmarkBytesFormat/bytefmt_test.go:1493/float_precision_1_format_%_.1f_1-8                      250728          4580 ns/op
BenchmarkBytesFormat/bytefmt_test.go:1548/float_precision_1_format_%_.1f_1024-8                   255877          4571 ns/op
BenchmarkBytesFormat/bytefmt_test.go:1603/float_precision_1_format_%_.1f_1128-8                   246607          4778 ns/op
BenchmarkBytesFormat/bytefmt_test.go:1658/integer_format_%_d_0-8                                  377036          2998 ns/op
BenchmarkBytesFormat/bytefmt_test.go:1713/integer_format_%_d_1-8                                  402634          2959 ns/op
BenchmarkBytesFormat/bytefmt_test.go:1768/integer_format_%_d_1024-8                               402206          3020 ns/op
BenchmarkBytesFormat/bytefmt_test.go:1823/integer_format_%_d_1128-8                               410911          2983 ns/op
BenchmarkBytesFormat/bytefmt_test.go:1878/integer_precision_2_format_%_.2d_0-8                    374540          3265 ns/op
BenchmarkBytesFormat/bytefmt_test.go:1951/integer_precision_2_format_%_.2d_1-8                    361249          3342 ns/op
BenchmarkBytesFormat/bytefmt_test.go:2030/integer_precision_2_format_%_.2d_1024-8                 379489          3349 ns/op
BenchmarkBytesFormat/bytefmt_test.go:2109/integer_precision_2_format_%_.2d_1128-8                 376008          3264 ns/op
BenchmarkBytesFormat/bytefmt_test.go:2189/name_%_d_1124-8                                         352525          3380 ns/op
PASS
ok      github.com/pfmt/bytefmt 53.576s
```
