// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bytefmt_test

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"testing"

	"github.com/pfmt/bytefmt"
)

var namesTests = []struct {
	line  string
	names []string
	want  []string
	bench bool
}{
	{
		line:  testline(),
		names: []string{"B", "K", "M", "G", "T", "P", "E"},
		want:  []string{"B", "K", "M", "G", "T", "P", "E"},
		bench: true,
	}, {
		line:  testline(),
		names: []string{},
		want:  []string{"B", "K", "M", "G", "T", "P", "E"},
	}, {
		line:  testline(),
		names: []string{"B"},
		want:  []string{"B", "K", "M", "G", "T", "P", "E"},
	}, {
		line:  testline(),
		names: []string{"B", "K"},
		want:  []string{"B", "K", "M", "G", "T", "P", "E"},
	}, {
		line:  testline(),
		names: []string{"B", "Kilobyte"},
		want:  []string{"B", "Kilobyte", "M", "G", "T", "P", "E"},
	}, {
		line:  testline(),
		names: []string{"B", "K", "M", "G", "T", "P", "E", "Zettabyte"},
		want:  []string{"B", "K", "M", "G", "T", "P", "E", "Zettabyte"},
	},
}

func TestNamesInitialize(t *testing.T) {
	for _, tt := range namesTests {
		tt := tt

		t.Run(tt.line+"/"+strings.Join(tt.names, " "), func(t *testing.T) {
			t.Parallel()

			b := bytefmt.New(0, tt.names...)
			got := b.Names()
			if strings.Join(got, "") != strings.Join(tt.want, "") {
				t.Errorf("\nwant name: %#v\n got name: %#v\ntest: %s", tt.want, got, tt.line)
			}
		})
	}
}

func TestNamesUpdate(t *testing.T) {
	for _, tt := range namesTests {
		tt := tt

		t.Run(tt.line+"/"+strings.Join(tt.names, " "), func(t *testing.T) {
			t.Parallel()

			b := bytefmt.New(0)
			got := b.Names(tt.names...)
			if strings.Join(got, "") != strings.Join(tt.want, "") {
				t.Errorf("\nwant name: %#v\n got name %#v\ntest: %s", tt.want, got, tt.line)
			}
		})
	}
}

func BenchmarkNamesInitialize(b *testing.B) {
	b.ReportAllocs()

	for _, tt := range namesTests {
		if !tt.bench {
			continue
		}

		b.Run(tt.line+"/"+strings.Join(tt.names, " "), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = bytefmt.New(0, tt.names...)
			}
		})
	}
}

func BenchmarkNamesUpdate(b *testing.B) {
	b.ReportAllocs()
	for _, tt := range namesTests {
		if !tt.bench {
			continue
		}

		b.Run(tt.line+"/"+strings.Join(tt.names, " "), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b := bytefmt.New(0)
				_ = b.Names(tt.names...)
			}
		})
	}
}

var bytesKilobytesTests = []struct {
	name  string
	line  string
	bytes uint64
	want  float64
	bench bool
}{
	{
		name:  "exactly one kilobyte",
		line:  testline(),
		bytes: 1024,
		want:  1,
	}, {
		name:  "less than one kilobyte",
		line:  testline(),
		bytes: 42,
		want:  0.041015625,
	}, {
		name:  "more than one kilobyte",
		line:  testline(),
		bytes: 100500,
		want:  98.14453125,
		bench: true,
	},
}

func TestBytesKilobytes(t *testing.T) {
	for _, tt := range bytesKilobytesTests {
		tt := tt

		t.Run(tt.line+"/"+tt.name+strconv.FormatUint(tt.bytes, 10), func(t *testing.T) {
			t.Parallel()

			got := bytefmt.Bytes{Value: tt.bytes}.Kilobytes()
			if got != tt.want {
				t.Errorf("\nwant kilobytes: %v\n got expected: %v\ntest: %s", tt.want, got, tt.line)
			}
		})
	}
}

func BenchmarkBytesKilobytes(b *testing.B) {
	b.ReportAllocs()

	for _, tt := range bytesKilobytesTests {
		if !tt.bench {
			continue
		}

		b.Run(tt.line+"/"+tt.name+strconv.FormatUint(tt.bytes, 10), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = bytefmt.Bytes{Value: tt.bytes}.Kilobytes()
			}
		})
	}
}

var bytesMegabytesTests = []struct {
	name  string
	line  string
	bytes uint64
	want  float64
	bench bool
}{
	{
		name:  "exactly one megabyte",
		line:  testline(),
		bytes: 1048576,
		want:  1,
	}, {
		name:  "less than one megabyte",
		line:  testline(),
		bytes: 42,
		want:  4.00543212890625e-05,
	}, {
		name:  "more than one megabyte",
		line:  testline(),
		bytes: 10050000,
		want:  9.584426879882812,
		bench: true,
	},
}

func TestBytesMegabytes(t *testing.T) {
	for _, tt := range bytesMegabytesTests {
		tt := tt

		t.Run(tt.line+"/"+tt.name+strconv.FormatUint(tt.bytes, 10), func(t *testing.T) {
			t.Parallel()

			got := bytefmt.Bytes{Value: tt.bytes}.Megabytes()
			if got != tt.want {
				t.Errorf("\nwant megabytes: %v\n got megabytes: %v\ntest: %s", tt.want, got, tt.line)
			}
		})
	}
}

func BenchmarkBytesMegabytes(b *testing.B) {
	b.ReportAllocs()

	for _, tt := range bytesMegabytesTests {
		if !tt.bench {
			continue
		}

		b.Run(tt.line+"/"+tt.name+strconv.FormatUint(tt.bytes, 10), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = bytefmt.Bytes{Value: tt.bytes}.Megabytes()
			}
		})
	}
}

var bytesGigabytesTests = []struct {
	name  string
	line  string
	bytes uint64
	want  float64
	bench bool
}{
	{
		name:  "exactly one gigabyte",
		line:  testline(),
		bytes: 1073741824,
		want:  1,
	}, {
		name:  "less than one gigabyte",
		line:  testline(),
		bytes: 42,
		want:  3.91155481338501e-08,
	}, {
		name:  "more than one gigabyte",
		line:  testline(),
		bytes: 10050000000,
		want:  9.359791874885559,
		bench: true,
	},
}

func TestBytesGigabytes(t *testing.T) {
	for _, tt := range bytesGigabytesTests {
		tt := tt
		t.Run(tt.line+"/"+tt.name+strconv.FormatUint(tt.bytes, 10), func(t *testing.T) {
			t.Parallel()

			got := bytefmt.Bytes{Value: tt.bytes}.Gigabytes()
			if got != tt.want {
				t.Errorf("\nwant gigabytes: %v\n got gigabytes: %v\ntest: %s", tt.want, got, tt.line)
			}
		})
	}
}

func BenchmarkBytesGigabytes(b *testing.B) {
	b.ReportAllocs()

	for _, tt := range bytesGigabytesTests {
		if !tt.bench {
			continue
		}

		b.Run(tt.line+"/"+tt.name+strconv.FormatUint(tt.bytes, 10), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = bytefmt.Bytes{Value: tt.bytes}.Gigabytes()
			}
		})
	}
}

var bytesTerabytesTests = []struct {
	name  string
	line  string
	bytes uint64
	want  float64
	bench bool
}{
	{
		name:  "exactly one terabyte",
		line:  testline(),
		bytes: 1099511627776,
		want:  1,
	}, {
		name:  "less than one terabyte",
		line:  testline(),
		bytes: 42,
		want:  3.8198777474462986e-11,
	}, {
		name:  "more than one terabyte",
		line:  testline(),
		bytes: 10050000000000,
		want:  9.140421752817929,
		bench: true,
	},
}

func TestBytesTerabytes(t *testing.T) {
	for _, tt := range bytesTerabytesTests {
		tt := tt
		t.Run(tt.line+"/"+tt.name+strconv.FormatUint(tt.bytes, 10), func(t *testing.T) {
			t.Parallel()

			got := bytefmt.Bytes{Value: tt.bytes}.Terabytes()
			if got != tt.want {
				t.Errorf("\nwant terabytes: %v\n got terabytes: %v\ntest: %s", tt.want, got, tt.line)
			}
		})
	}
}

func BenchmarkBytesTerabytes(b *testing.B) {
	b.ReportAllocs()

	for _, tt := range bytesTerabytesTests {
		if !tt.bench {
			continue
		}

		b.Run(tt.line+"/"+tt.name+strconv.FormatUint(tt.bytes, 10), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = bytefmt.Bytes{Value: tt.bytes}.Terabytes()
			}
		})
	}
}

var bytesPetabytesTests = []struct {
	name  string
	line  string
	bytes uint64
	want  float64
	bench bool
}{
	{
		name:  "exactly one petabyte",
		line:  testline(),
		bytes: 1125899906842624,
		want:  1,
	}, {
		name:  "less than one petabyte",
		line:  testline(),
		bytes: 42,
		want:  3.730349362740526e-14,
	}, {
		name:  "more than one petabyte",
		line:  testline(),
		bytes: 10050000000000000,
		want:  8.926193117986259,
		bench: true,
	},
}

func TestBytesPetabytes(t *testing.T) {
	for _, tt := range bytesPetabytesTests {
		tt := tt
		t.Run(tt.line+"/"+tt.name+strconv.FormatUint(tt.bytes, 10), func(t *testing.T) {
			t.Parallel()

			got := bytefmt.Bytes{Value: tt.bytes}.Petabytes()
			if got != tt.want {
				t.Errorf("\nwant petabytes: %v\n got petabytes: %v\ntest: %s", tt.want, got, tt.line)
			}
		})
	}
}

func BenchmarkBytesPetabytes(b *testing.B) {
	b.ReportAllocs()

	for _, tt := range bytesPetabytesTests {
		if !tt.bench {
			continue
		}

		b.Run(tt.line+"/"+tt.name+strconv.FormatUint(tt.bytes, 10), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = bytefmt.Bytes{Value: tt.bytes}.Petabytes()
			}
		})
	}
}

var BytesExabytesTests = []struct {
	name  string
	line  string
	bytes uint64
	want  float64
	bench bool
}{
	{
		name:  "exactly one exabyte",
		line:  testline(),
		bytes: 1152921504606846976,
		want:  1,
	}, {
		name:  "less than one exabyte",
		line:  testline(),
		bytes: 42,
		want:  3.642919299551295e-17,
	}, {
		name:  "more than one exabyte",
		line:  testline(),
		bytes: 10050000000000000000,
		want:  8.716985466783456,
		bench: true,
	},
}

func TestBytesExabytes(t *testing.T) {
	for _, tt := range BytesExabytesTests {
		tt := tt
		t.Run(tt.line+"/"+tt.name+strconv.FormatUint(tt.bytes, 10), func(t *testing.T) {
			t.Parallel()

			got := bytefmt.Bytes{Value: tt.bytes}.Exabytes()
			if got != tt.want {
				t.Errorf("\nwant exabytes: %v\n got exabytes: %v\ntest: %s", tt.want, got, tt.line)
			}
		})
	}
}

func BenchmarkBytesExabytes(b *testing.B) {
	b.ReportAllocs()

	for _, tt := range BytesExabytesTests {
		if !tt.bench {
			continue
		}

		b.Run(tt.line+"/"+tt.name+strconv.FormatUint(tt.bytes, 10), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = bytefmt.Bytes{Value: tt.bytes}.Exabytes()
			}
		})
	}
}

var bytesStringTests = []struct {
	name  string
	line  string
	bytes uint64
	want  string
	bench bool
}{
	{
		name:  "zero byte",
		line:  testline(),
		bytes: 0,
		want:  "0B",
	}, {
		name:  "less than one kilobyte",
		line:  testline(),
		bytes: 1,
		want:  "1B",
	}, {
		name:  "exactly one kilobyte",
		line:  testline(),
		bytes: 1024,
		want:  "1K",
	}, {
		name:  "more than one kilobyte",
		line:  testline(),
		bytes: 1128,
		want:  "1.1015625K",
		bench: true,
	},
}

func TestBytesString(t *testing.T) {
	for _, tt := range bytesStringTests {
		tt := tt
		t.Run(tt.line+"/"+tt.name+strconv.FormatUint(tt.bytes, 10), func(t *testing.T) {
			t.Parallel()

			got := bytefmt.New(tt.bytes).String()
			if got != tt.want {
				t.Errorf("\nwant string: %#v\n got string: %#v\ntest: %s", tt.want, got, tt.line)
			}
		})
	}
}

func BenchmarkBytesString(b *testing.B) {
	b.ReportAllocs()

	for _, tt := range bytesStringTests {
		if !tt.bench {
			continue
		}

		b.Run(tt.line+"/"+tt.name+strconv.FormatUint(tt.bytes, 10), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = bytefmt.New(tt.bytes).String()
			}
		})
	}
}

var bytesFormatTestc = []struct {
	name   string
	line   string
	bytes  uint64
	format string
	names  []string
	want   string
	bench  bool
}{
	{
		name:   "general format",
		line:   testline(),
		bytes:  0,
		format: "%v",
		want:   "0B",
	}, {
		name:   "general format",
		line:   testline(),
		bytes:  0,
		format: "% v",
		want:   "0 B",
		bench:  true,
	}, {
		name:   "general format",
		line:   testline(),
		bytes:  0,
		format: "% 2v",
		want:   "0  B",
	}, {
		name:   "general format",
		line:   testline(),
		bytes:  0,
		format: "%3v",
		want:   " 0B",
	}, {
		name:   "general format",
		line:   testline(),
		bytes:  0,
		format: "%03v",
		want:   "00B",
	}, {
		name:   "general format",
		line:   testline(),
		bytes:  0,
		format: "%-3v",
		want:   "0B ",
	}, {
		name:   "general format",
		line:   testline(),
		bytes:  0,
		format: "%-03v",
		want:   "0B ",
	}, {
		name:   "general format",
		line:   testline(),
		bytes:  0,
		format: "%+v",
		want:   "0B",
	}, {
		name:   "general format",
		line:   testline(),
		bytes:  0,
		format: "%#v",
		want:   "0B",
	},

	{
		name:   "general format",
		line:   testline(),
		bytes:  1,
		format: "%v",
		want:   "1B",
	}, {
		name:   "general format",
		line:   testline(),
		bytes:  1,
		format: "% v",
		want:   "1 B",
		bench:  true,
	}, {
		name:   "general format",
		line:   testline(),
		bytes:  1,
		format: "% 2v",
		want:   "1  B",
	}, {
		name:   "general format",
		line:   testline(),
		bytes:  1,
		format: "%3v",
		want:   " 1B",
	}, {
		name:   "general format",
		line:   testline(),
		bytes:  1,
		format: "%03v",
		want:   "01B",
	}, {
		name:   "general format",
		line:   testline(),
		bytes:  1,
		format: "%-3v",
		want:   "1B ",
	}, {
		name:   "general format",
		line:   testline(),
		bytes:  1,
		format: "%-03v",
		want:   "1B ",
	}, {
		name:   "general format",
		line:   testline(),
		bytes:  1,
		format: "%+v",
		want:   "1B",
	}, {
		name:   "general format",
		line:   testline(),
		bytes:  1,
		format: "%#v",
		want:   "1B",
	},

	{
		name:   "general format",
		line:   testline(),
		bytes:  1024,
		format: "%v",
		want:   "1K",
	}, {
		name:   "general format",
		line:   testline(),
		bytes:  1024,
		format: "% v",
		want:   "1 K",
		bench:  true,
	}, {
		name:   "general format",
		line:   testline(),
		bytes:  1024,
		format: "% 2v",
		want:   "1  K",
	}, {
		name:   "general format",
		line:   testline(),
		bytes:  1024,
		format: "%3v",
		want:   " 1K",
	}, {
		name:   "general format",
		line:   testline(),
		bytes:  1024,
		format: "%03v",
		want:   "01K",
	}, {
		name:   "general format",
		line:   testline(),
		bytes:  1024,
		format: "%-3v",
		want:   "1K ",
	}, {
		name:   "general format",
		line:   testline(),
		bytes:  1024,
		format: "%-03v",
		want:   "1K ",
	}, {
		name:   "general format",
		line:   testline(),
		bytes:  1024,
		format: "%+v",
		want:   "1K",
	}, {
		name:   "general format",
		line:   testline(),
		bytes:  1024,
		format: "%#v",
		want:   "1K",
	},

	{
		name:   "general format",
		line:   testline(),
		bytes:  1128,
		format: "%v",
		want:   "1.1015625K",
	}, {
		name:   "general format",
		line:   testline(),
		bytes:  1128,
		format: "% v",
		want:   "1.1015625 K",
		bench:  true,
	}, {
		name:   "general format",
		line:   testline(),
		bytes:  1128,
		format: "% 2v",
		want:   "1.1015625  K",
	}, {
		name:   "general format",
		line:   testline(),
		bytes:  1128,
		format: "%11v",
		want:   " 1.1015625K",
	}, {
		name:   "general format",
		line:   testline(),
		bytes:  1128,
		format: "%011v",
		want:   "01.1015625K",
	}, {
		name:   "general format",
		line:   testline(),
		bytes:  1128,
		format: "%-11v",
		want:   "1.1015625K ",
	}, {
		name:   "general format",
		line:   testline(),
		bytes:  1128,
		format: "%-011v",
		want:   "1.1015625K ",
	}, {
		name:   "general format",
		line:   testline(),
		bytes:  1128,
		format: "%+v",
		want:   "1.1015625K",
	}, {
		name:   "general format",
		line:   testline(),
		bytes:  1128,
		format: "%#v",
		want:   "1.1015625K",
	}, {
		name:   "string format",
		line:   testline(),
		bytes:  0,
		format: "%s",
		want:   "0B",
	}, {
		name:   "string format",
		line:   testline(),
		bytes:  0,
		format: "% s",
		want:   "0 B",
		bench:  true,
	}, {
		name:   "string format",
		line:   testline(),
		bytes:  0,
		format: "% 2s",
		want:   "0  B",
	}, {
		name:   "string format",
		line:   testline(),
		bytes:  0,
		format: "%3s",
		want:   " 0B",
	}, {
		name:   "string format",
		line:   testline(),
		bytes:  0,
		format: "%03s",
		want:   "00B",
	}, {
		name:   "string format",
		line:   testline(),
		bytes:  0,
		format: "%-3s",
		want:   "0B ",
	}, {
		name:   "string format",
		line:   testline(),
		bytes:  0,
		format: "%-03s",
		want:   "0B ",
	}, {
		name:   "string format",
		line:   testline(),
		bytes:  0,
		format: "%+s",
		want:   "0B",
	}, {
		name:   "string format",
		line:   testline(),
		bytes:  0,
		format: "%#s",
		want:   "0B",
	},

	{
		name:   "string format",
		line:   testline(),
		bytes:  1,
		format: "%s",
		want:   "1B",
	}, {
		name:   "string format",
		line:   testline(),
		bytes:  1,
		format: "% s",
		want:   "1 B",
		bench:  true,
	}, {
		name:   "string format",
		line:   testline(),
		bytes:  1,
		format: "% 2s",
		want:   "1  B",
	}, {
		name:   "string format",
		line:   testline(),
		bytes:  1,
		format: "%3s",
		want:   " 1B",
	}, {
		name:   "string format",
		line:   testline(),
		bytes:  1,
		format: "%03s",
		want:   "01B",
	}, {
		name:   "string format",
		line:   testline(),
		bytes:  1,
		format: "%-3s",
		want:   "1B ",
	}, {
		name:   "string format",
		line:   testline(),
		bytes:  1,
		format: "%-03s",
		want:   "1B ",
	}, {
		name:   "string format",
		line:   testline(),
		bytes:  1,
		format: "%+s",
		want:   "1B",
	}, {
		name:   "string format",
		line:   testline(),
		bytes:  1,
		format: "%#s",
		want:   "1B",
	},

	{
		name:   "string format",
		line:   testline(),
		bytes:  1024,
		format: "%s",
		want:   "1K",
	}, {
		name:   "string format",
		line:   testline(),
		bytes:  1024,
		format: "% s",
		want:   "1 K",
		bench:  true,
	}, {
		name:   "string format",
		line:   testline(),
		bytes:  1024,
		format: "% 2s",
		want:   "1  K",
	}, {
		name:   "string format",
		line:   testline(),
		bytes:  1024,
		format: "%3s",
		want:   " 1K",
	}, {
		name:   "string format",
		line:   testline(),
		bytes:  1024,
		format: "%03s",
		want:   "01K",
	}, {
		name:   "string format",
		line:   testline(),
		bytes:  1024,
		format: "%-3s",
		want:   "1K ",
	}, {
		name:   "string format",
		line:   testline(),
		bytes:  1024,
		format: "%-03s",
		want:   "1K ",
	}, {
		name:   "string format",
		line:   testline(),
		bytes:  1024,
		format: "%+s",
		want:   "1K",
	}, {
		name:   "string format",
		line:   testline(),
		bytes:  1024,
		format: "%#s",
		want:   "1K",
	},

	{
		name:   "string format",
		line:   testline(),
		bytes:  1128,
		format: "%s",
		want:   "1.1015625K",
	}, {
		name:   "string format",
		line:   testline(),
		bytes:  1128,
		format: "% s",
		want:   "1.1015625 K",
		bench:  true,
	}, {
		name:   "string format",
		line:   testline(),
		bytes:  1128,
		format: "% 2s",
		want:   "1.1015625  K",
	}, {
		name:   "string format",
		line:   testline(),
		bytes:  1128,
		format: "%11s",
		want:   " 1.1015625K",
	}, {
		name:   "string format",
		line:   testline(),
		bytes:  1128,
		format: "%011s",
		want:   "01.1015625K",
	}, {
		name:   "string format",
		line:   testline(),
		bytes:  1128,
		format: "%-11s",
		want:   "1.1015625K ",
	}, {
		name:   "string format",
		line:   testline(),
		bytes:  1128,
		format: "%-011s",
		want:   "1.1015625K ",
	}, {
		name:   "string format",
		line:   testline(),
		bytes:  1128,
		format: "%+s",
		want:   "1.1015625K",
	}, {
		name:   "string format",
		line:   testline(),
		bytes:  1128,
		format: "%#s",
		want:   "1.1015625K",
	}, {
		name:   "double-quoted string format",
		line:   testline(),
		bytes:  0,
		format: "%q",
		want:   `"0B"`,
	}, {
		name:   "double-quoted string format",
		line:   testline(),
		bytes:  0,
		format: "% q",
		want:   `"0 B"`,
		bench:  true,
	}, {
		name:   "double-quoted string format",
		line:   testline(),
		bytes:  0,
		format: "% 2q",
		want:   `"0  B"`,
	}, {
		name:   "double-quoted string format",
		line:   testline(),
		bytes:  0,
		format: "%3q",
		want:   `" 0B"`,
	}, {
		name:   "double-quoted string format",
		line:   testline(),
		bytes:  0,
		format: "%03q",
		want:   `"00B"`,
	}, {
		name:   "double-quoted string format",
		line:   testline(),
		bytes:  0,
		format: "%-3q",
		want:   `"0B "`,
	}, {
		name:   "double-quoted string format",
		line:   testline(),
		bytes:  0,
		format: "%-03q",
		want:   `"0B "`,
	}, {
		name:   "double-quoted string format",
		line:   testline(),
		bytes:  0,
		format: "%+q",
		want:   `"0B"`,
	}, {
		name:   "double-quoted string format",
		line:   testline(),
		bytes:  0,
		format: "%#q",
		want:   "`0B`",
	},

	{
		name:   "double-quoted string format",
		line:   testline(),
		bytes:  1,
		format: "%q",
		want:   `"1B"`,
	}, {
		name:   "double-quoted string format",
		line:   testline(),
		bytes:  1,
		format: "% q",
		want:   `"1 B"`,
		bench:  true,
	}, {
		name:   "double-quoted string format",
		line:   testline(),
		bytes:  1,
		format: "% 2q",
		want:   `"1  B"`,
	}, {
		name:   "double-quoted string format",
		line:   testline(),
		bytes:  1,
		format: "%3q",
		want:   `" 1B"`,
	}, {
		name:   "double-quoted string format",
		line:   testline(),
		bytes:  1,
		format: "%03q",
		want:   `"01B"`,
	}, {
		name:   "double-quoted string format",
		line:   testline(),
		bytes:  1,
		format: "%-3q",
		want:   `"1B "`,
	}, {
		name:   "double-quoted string format",
		line:   testline(),
		bytes:  1,
		format: "%-03q",
		want:   `"1B "`,
	}, {
		name:   "double-quoted string format",
		line:   testline(),
		bytes:  1,
		format: "%+q",
		want:   `"1B"`,
	}, {
		name:   "double-quoted string format",
		line:   testline(),
		bytes:  1,
		format: "%#q",
		want:   "`1B`",
	},

	{
		name:   "double-quoted string format",
		line:   testline(),
		bytes:  1024,
		format: "%q",
		want:   `"1K"`,
	}, {
		name:   "double-quoted string format",
		line:   testline(),
		bytes:  1024,
		format: "% q",
		want:   `"1 K"`,
		bench:  true,
	}, {
		name:   "double-quoted string format",
		line:   testline(),
		bytes:  1024,
		format: "% 2q",
		want:   `"1  K"`,
	}, {
		name:   "double-quoted string format",
		line:   testline(),
		bytes:  1024,
		format: "%3q",
		want:   `" 1K"`,
	}, {
		name:   "double-quoted string format",
		line:   testline(),
		bytes:  1024,
		format: "%03q",
		want:   `"01K"`,
	}, {
		name:   "double-quoted string format",
		line:   testline(),
		bytes:  1024,
		format: "%-3q",
		want:   `"1K "`,
	}, {
		name:   "double-quoted string format",
		line:   testline(),
		bytes:  1024,
		format: "%-03q",
		want:   `"1K "`,
	}, {
		name:   "double-quoted string format",
		line:   testline(),
		bytes:  1024,
		format: "%+q",
		want:   `"1K"`,
	}, {
		name:   "double-quoted string format",
		line:   testline(),
		bytes:  1024,
		format: "%#q",
		want:   "`1K`",
	},

	{
		name:   "double-quoted string format",
		line:   testline(),
		bytes:  1128,
		format: "%q",
		want:   `"1.1015625K"`,
	}, {
		name:   "double-quoted string format",
		line:   testline(),
		bytes:  1128,
		format: "% q",
		want:   `"1.1015625 K"`,
		bench:  true,
	}, {
		name:   "double-quoted string format",
		line:   testline(),
		bytes:  1128,
		format: "% 2q",
		want:   `"1.1015625  K"`,
	}, {
		name:   "double-quoted string format",
		line:   testline(),
		bytes:  1128,
		format: "%11q",
		want:   `" 1.1015625K"`,
	}, {
		name:   "double-quoted string format",
		line:   testline(),
		bytes:  1128,
		format: "%011q",
		want:   `"01.1015625K"`,
	}, {
		name:   "double-quoted string format",
		line:   testline(),
		bytes:  1128,
		format: "%-11q",
		want:   `"1.1015625K "`,
	}, {
		name:   "double-quoted string format",
		line:   testline(),
		bytes:  1128,
		format: "%-011q",
		want:   `"1.1015625K "`,
	}, {
		name:   "double-quoted string format",
		line:   testline(),
		bytes:  1128,
		format: "%+q",
		want:   `"1.1015625K"`,
	}, {
		name:   "double-quoted string format",
		line:   testline(),
		bytes:  1128,
		format: "%#q",
		want:   "`1.1015625K`",
	},

	{
		name:   "float format",
		line:   testline(),
		bytes:  0,
		format: "%f",
		want:   "0.000000B",
	}, {
		name:   "float format",
		line:   testline(),
		bytes:  0,
		format: "% f",
		want:   "0.000000 B",
		bench:  true,
	}, {
		name:   "float format",
		line:   testline(),
		bytes:  0,
		format: "% 2f",
		want:   "0.000000  B",
	}, {
		name:   "float format",
		line:   testline(),
		bytes:  0,
		format: "%10f",
		want:   " 0.000000B",
	}, {
		name:   "float format",
		line:   testline(),
		bytes:  0,
		format: "%010f",
		want:   "00.000000B",
	}, {
		name:   "float format",
		line:   testline(),
		bytes:  0,
		format: "%-10f",
		want:   "0.000000B ",
	}, {
		name:   "float format",
		line:   testline(),
		bytes:  0,
		format: "%-010f",
		want:   "0.000000B ",
	}, {
		name:   "float format",
		line:   testline(),
		bytes:  0,
		format: "%+f",
		want:   "+0.000000B",
	}, {
		name:   "float format",
		line:   testline(),
		bytes:  0,
		format: "%#f",
		want:   "0.000000B",
	}, {
		name:   "float format",
		line:   testline(),
		bytes:  1,
		format: "%f",
		want:   "1.000000B",
		bench:  true,
	}, {
		name:   "float format",
		line:   testline(),
		bytes:  1,
		format: "% 2f",
		want:   "1.000000  B",
	}, {
		name:   "float format",
		line:   testline(),
		bytes:  1,
		format: "%10f",
		want:   " 1.000000B",
	}, {
		name:   "float format",
		line:   testline(),
		bytes:  1,
		format: "%010f",
		want:   "01.000000B",
	}, {
		name:   "float format",
		line:   testline(),
		bytes:  1,
		format: "%-10f",
		want:   "1.000000B ",
	}, {
		name:   "float format",
		line:   testline(),
		bytes:  1,
		format: "%-010f",
		want:   "1.000000B ",
	}, {
		name:   "float format",
		line:   testline(),
		bytes:  1,
		format: "%+f",
		want:   "+1.000000B",
	}, {
		name:   "float format",
		line:   testline(),
		bytes:  1,
		format: "%#f",
		want:   "1.000000B",
	}, {
		name:   "float format",
		line:   testline(),
		bytes:  1024,
		format: "%f",
		want:   "1.000000K",
	}, {
		name:   "float format",
		line:   testline(),
		bytes:  1024,
		format: "% f",
		want:   "1.000000 K",
		bench:  true,
	}, {
		name:   "float format",
		line:   testline(),
		bytes:  1024,
		format: "% 2f",
		want:   "1.000000  K",
	}, {
		name:   "float format",
		line:   testline(),
		bytes:  1024,
		format: "%10f",
		want:   " 1.000000K",
	}, {
		name:   "float format",
		line:   testline(),
		bytes:  1024,
		format: "%010f",
		want:   "01.000000K",
	}, {
		name:   "float format",
		line:   testline(),
		bytes:  1024,
		format: "%-10f",
		want:   "1.000000K ",
	}, {
		name:   "float format",
		line:   testline(),
		bytes:  1024,
		format: "%-010f",
		want:   "1.000000K ",
	}, {
		name:   "float format",
		line:   testline(),
		bytes:  1024,
		format: "%+f",
		want:   "+1.000000K",
	}, {
		name:   "float format",
		line:   testline(),
		bytes:  1024,
		format: "%#f",
		want:   "1.000000K",
	}, {
		name:   "float format",
		line:   testline(),
		bytes:  1128,
		format: "%f",
		want:   "1.101562K",
	}, {
		name:   "float format",
		line:   testline(),
		bytes:  1128,
		format: "% f",
		want:   "1.101562 K",
		bench:  true,
	}, {
		name:   "float format",
		line:   testline(),
		bytes:  1128,
		format: "% 2f",
		want:   "1.101562  K",
	}, {
		name:   "float format",
		line:   testline(),
		bytes:  1128,
		format: "%10f",
		want:   " 1.101562K",
	}, {
		name:   "float format",
		line:   testline(),
		bytes:  1128,
		format: "%010f",
		want:   "01.101562K",
	}, {
		name:   "float format",
		line:   testline(),
		bytes:  1128,
		format: "%-10f",
		want:   "1.101562K ",
	}, {
		name:   "float format",
		line:   testline(),
		bytes:  1128,
		format: "%-010f",
		want:   "1.101562K ",
	}, {
		name:   "float format",
		line:   testline(),
		bytes:  1128,
		format: "%+f",
		want:   "+1.101562K",
	}, {
		name:   "float format",
		line:   testline(),
		bytes:  1128,
		format: "%#f",
		want:   "1.101562K",
	}, {
		name:   "float precision 1 format",
		line:   testline(),
		bytes:  0,
		format: "%.1f",
		want:   "0.0B",
	}, {
		name:   "float precision 1 format",
		line:   testline(),
		bytes:  0,
		format: "% .1f",
		want:   "0.0 B",
		bench:  true,
	}, {
		name:   "float precision 1 format",
		line:   testline(),
		bytes:  0,
		format: "% 2.1f",
		want:   "0.0  B",
	}, {
		name:   "float precision 1 format",
		line:   testline(),
		bytes:  0,
		format: "%4.1f",
		want:   "0.0B",
	}, {
		name:   "float precision 1 format",
		line:   testline(),
		bytes:  0,
		format: "%6.1f",
		want:   "  0.0B",
	}, {
		name:   "float precision 1 format",
		line:   testline(),
		bytes:  0,
		format: "%04.1f",
		want:   "0.0B",
	}, {
		name:   "float precision 1 format",
		line:   testline(),
		bytes:  0,
		format: "%06.1f",
		want:   "000.0B",
	}, {
		name:   "float precision 1 format",
		line:   testline(),
		bytes:  0,
		format: "%-5.1f",
		want:   "0.0B ",
	}, {
		name:   "float precision 1 format",
		line:   testline(),
		bytes:  0,
		format: "%-05.1f",
		want:   "0.0B ",
	}, {
		name:   "float precision 1 format",
		line:   testline(),
		bytes:  0,
		format: "%+.1f",
		want:   "+0.0B",
	}, {
		name:   "float precision 1 format",
		line:   testline(),
		bytes:  0,
		format: "%#.1f",
		want:   "0.0B",
	}, {
		name:   "float precision 1 format",
		line:   testline(),
		bytes:  1,
		format: "%.1f",
		want:   "1.0B",
	}, {
		name:   "float precision 1 format",
		line:   testline(),
		bytes:  1,
		format: "% .1f",
		want:   "1.0 B",
		bench:  true,
	}, {
		name:   "float precision 1 format",
		line:   testline(),
		bytes:  1,
		format: "% 2.1f",
		want:   "1.0  B",
	}, {
		name:   "float precision 1 format",
		line:   testline(),
		bytes:  1,
		format: "%5.1f",
		want:   " 1.0B",
	}, {
		name:   "float precision 1 format",
		line:   testline(),
		bytes:  1,
		format: "%05.1f",
		want:   "01.0B",
	}, {
		name:   "float precision 1 format",
		line:   testline(),
		bytes:  1,
		format: "%-5.1f",
		want:   "1.0B ",
	}, {
		name:   "float precision 1 format",
		line:   testline(),
		bytes:  1,
		format: "%-05.1f",
		want:   "1.0B ",
	}, {
		name:   "float precision 1 format",
		line:   testline(),
		bytes:  1,
		format: "%+.1f",
		want:   "+1.0B",
	}, {
		name:   "float precision 1 format",
		line:   testline(),
		bytes:  1,
		format: "%#.1f",
		want:   "1.0B",
	}, {
		name:   "float precision 1 format",
		line:   testline(),
		bytes:  1024,
		format: "%.1f",
		want:   "1.0K",
	}, {
		name:   "float precision 1 format",
		line:   testline(),
		bytes:  1024,
		format: "% .1f",
		want:   "1.0 K",
		bench:  true,
	}, {
		name:   "float precision 1 format",
		line:   testline(),
		bytes:  1024,
		format: "% 2.1f",
		want:   "1.0  K",
	}, {
		name:   "float precision 1 format",
		line:   testline(),
		bytes:  1024,
		format: "%5.1f",
		want:   " 1.0K",
	}, {
		name:   "float precision 1 format",
		line:   testline(),
		bytes:  1024,
		format: "%05.1f",
		want:   "01.0K",
	}, {
		name:   "float precision 1 format",
		line:   testline(),
		bytes:  1024,
		format: "%-5.1f",
		want:   "1.0K ",
	}, {
		name:   "float precision 1 format",
		line:   testline(),
		bytes:  1024,
		format: "%-05.1f",
		want:   "1.0K ",
	}, {
		name:   "float precision 1 format",
		line:   testline(),
		bytes:  1024,
		format: "%+.1f",
		want:   "+1.0K",
	}, {
		name:   "float precision 1 format",
		line:   testline(),
		bytes:  1024,
		format: "%#.1f",
		want:   "1.0K",
	}, {
		name:   "float precision 1 format",
		line:   testline(),
		bytes:  1128,
		format: "%.1f",
		want:   "1.1K",
	}, {
		name:   "float precision 1 format",
		line:   testline(),
		bytes:  1128,
		format: "% .1f",
		want:   "1.1 K",
		bench:  true,
	}, {
		name:   "float precision 1 format",
		line:   testline(),
		bytes:  1128,
		format: "% 2.1f",
		want:   "1.1  K",
	}, {
		name:   "float precision 1 format",
		line:   testline(),
		bytes:  1128,
		format: "%5.1f",
		want:   " 1.1K",
	}, {
		name:   "float precision 1 format",
		line:   testline(),
		bytes:  1128,
		format: "%05.1f",
		want:   "01.1K",
	}, {
		name:   "float precision 1 format",
		line:   testline(),
		bytes:  1128,
		format: "%-5.1f",
		want:   "1.1K ",
	}, {
		name:   "float precision 1 format",
		line:   testline(),
		bytes:  1128,
		format: "%-05.1f",
		want:   "1.1K ",
	}, {
		name:   "float precision 1 format",
		line:   testline(),
		bytes:  1128,
		format: "%+.1f",
		want:   "+1.1K",
	}, {
		name:   "float precision 1 format",
		line:   testline(),
		bytes:  1128,
		format: "%#.1f",
		want:   "1.1K",
	}, {
		name:   "integer format",
		line:   testline(),
		bytes:  0,
		format: "%d",
		want:   "0B",
	}, {
		name:   "integer format",
		line:   testline(),
		bytes:  0,
		format: "% d",
		want:   "0 B",
		bench:  true,
	}, {
		name:   "integer format",
		line:   testline(),
		bytes:  0,
		format: "% 2d",
		want:   "0  B",
	}, {
		name:   "integer format",
		line:   testline(),
		bytes:  0,
		format: "%3d",
		want:   " 0B",
	}, {
		name:   "integer format",
		line:   testline(),
		bytes:  0,
		format: "%03d",
		want:   "00B",
	}, {
		name:   "integer format",
		line:   testline(),
		bytes:  0,
		format: "%-3d",
		want:   "0B ",
	}, {
		name:   "integer format",
		line:   testline(),
		bytes:  0,
		format: "%-03d",
		want:   "0B ",
	}, {
		name:   "integer format",
		line:   testline(),
		bytes:  0,
		format: "%+d",
		want:   "+0B",
	}, {
		name:   "integer format",
		line:   testline(),
		bytes:  0,
		format: "%#d",
		want:   "0B",
	}, {
		name:   "integer format",
		line:   testline(),
		bytes:  1,
		format: "%d",
		want:   "1B",
	}, {
		name:   "integer format",
		line:   testline(),
		bytes:  1,
		format: "% d",
		want:   "1 B",
		bench:  true,
	}, {
		name:   "integer format",
		line:   testline(),
		bytes:  1,
		format: "% 2d",
		want:   "1  B",
	}, {
		name:   "integer format",
		line:   testline(),
		bytes:  1,
		format: "%3d",
		want:   " 1B",
	}, {
		name:   "integer format",
		line:   testline(),
		bytes:  1,
		format: "%03d",
		want:   "01B",
	}, {
		name:   "integer format",
		line:   testline(),
		bytes:  1,
		format: "%-3d",
		want:   "1B ",
	}, {
		name:   "integer format",
		line:   testline(),
		bytes:  1,
		format: "%-03d",
		want:   "1B ",
	}, {
		name:   "integer format",
		line:   testline(),
		bytes:  1,
		format: "%+d",
		want:   "+1B",
	}, {
		name:   "integer format",
		line:   testline(),
		bytes:  1,
		format: "%#d",
		want:   "1B",
	}, {
		name:   "integer format",
		line:   testline(),
		bytes:  1024,
		format: "%d",
		want:   "1K",
	}, {
		name:   "integer format",
		line:   testline(),
		bytes:  1024,
		format: "% d",
		want:   "1 K",
		bench:  true,
	}, {
		name:   "integer format",
		line:   testline(),
		bytes:  1024,
		format: "% 2d",
		want:   "1  K",
	}, {
		name:   "integer format",
		line:   testline(),
		bytes:  1024,
		format: "%3d",
		want:   " 1K",
	}, {
		name:   "integer format",
		line:   testline(),
		bytes:  1024,
		format: "%03d",
		want:   "01K",
	}, {
		name:   "integer format",
		line:   testline(),
		bytes:  1024,
		format: "%-3d",
		want:   "1K ",
	}, {
		name:   "integer format",
		line:   testline(),
		bytes:  1024,
		format: "%-03d",
		want:   "1K ",
	}, {
		name:   "integer format",
		line:   testline(),
		bytes:  1024,
		format: "%+d",
		want:   "+1K",
	}, {
		name:   "integer format",
		line:   testline(),
		bytes:  1024,
		format: "%#d",
		want:   "1K",
	}, {
		name:   "integer format",
		line:   testline(),
		bytes:  1128,
		format: "%d",
		want:   "1K",
	}, {
		name:   "integer format",
		line:   testline(),
		bytes:  1128,
		format: "% d",
		want:   "1 K",
		bench:  true,
	}, {
		name:   "integer format",
		line:   testline(),
		bytes:  1128,
		format: "% 2d",
		want:   "1  K",
	}, {
		name:   "integer format",
		line:   testline(),
		bytes:  1128,
		format: "%3d",
		want:   " 1K",
	}, {
		name:   "integer format",
		line:   testline(),
		bytes:  1128,
		format: "%03d",
		want:   "01K",
	}, {
		name:   "integer format",
		line:   testline(),
		bytes:  1128,
		format: "%-3d",
		want:   "1K ",
	}, {
		name:   "integer format",
		line:   testline(),
		bytes:  1128,
		format: "%-03d",
		want:   "1K ",
	}, {
		name:   "integer format",
		line:   testline(),
		bytes:  1128,
		format: "%+d",
		want:   "+1K",
	}, {
		name:   "integer format",
		line:   testline(),
		bytes:  1128,
		format: "%#d",
		want:   "1K",
	}, {
		name:   "integer precision 2 format",
		line:   testline(),
		bytes:  0,
		format: "%.2d",
		want:   "00B",
	}, {
		name:   "integer precision 2 format",
		line:   testline(),
		bytes:  0,
		format: "% .2d",
		want:   "00 B",
		bench:  true,
	}, {
		name:   "integer precision 2 format",
		line:   testline(),
		bytes:  0,
		format: "% 2.2d",
		want:   "00  B",
	}, {
		name:   "integer precision 2 format",
		line:   testline(),
		bytes:  0,
		format: "%3.2d",
		want:   "00B",
	}, {
		name:   "integer precision 2 format",
		line:   testline(),
		bytes:  0,
		format: "%4.2d",
		want:   " 00B",
	}, {
		name:   "integer precision 2 format",
		line:   testline(),
		bytes:  0,
		format: "%03.2d",
		want:   "00B",
	}, {
		name:   "integer precision 2 format",
		line:   testline(),
		bytes:  0,
		format: "%-3.2d",
		want:   "00B",
	}, {
		name:   "integer precision 2 format",
		line:   testline(),
		bytes:  0,
		format: "%-4.2d",
		want:   "00B ",
	}, {
		name:   "integer precision 2 format",
		line:   testline(),
		bytes:  0,
		format: "%-03.2d",
		want:   "00B",
	}, {
		name:   "integer precision 2 format",
		line:   testline(),
		bytes:  0,
		format: "%-04.2d",
		want:   "00B ",
	}, {
		name:   "integer precision 2 format",
		line:   testline(),
		bytes:  0,
		format: "%+.2d",
		want:   "+00B",
	}, {
		name:   "integer precision 2 format",
		line:   testline(),
		bytes:  0,
		format: "%#.2d",
		want:   "00B",
	}, {
		name:   "integer precision 2 format",
		line:   testline(),
		bytes:  1,
		format: "%.2d",
		want:   "01B",
	}, {
		name:   "integer precision 2 format",
		line:   testline(),
		bytes:  1,
		format: "% .2d",
		want:   "01 B",
		bench:  true,
	}, {
		name:   "integer precision 2 format",
		line:   testline(),
		bytes:  1,
		format: "% 2.2d",
		want:   "01  B",
	}, {
		name:   "integer precision 2 format",
		line:   testline(),
		bytes:  1,
		format: "%2.2d",
		want:   "01B",
	}, {
		name:   "integer precision 2 format",
		line:   testline(),
		bytes:  1,
		format: "%4.2d",
		want:   " 01B",
	}, {
		name:   "integer precision 2 format",
		line:   testline(),
		bytes:  1,
		format: "%03.2d",
		want:   "01B",
	}, {
		name:   "integer precision 2 format",
		line:   testline(),
		bytes:  1,
		format: "%04.2d",
		want:   "001B",
	}, {
		name:   "integer precision 2 format",
		line:   testline(),
		bytes:  1,
		format: "%-3.2d",
		want:   "01B",
	}, {
		name:   "integer precision 2 format",
		line:   testline(),
		bytes:  1,
		format: "%-4.2d",
		want:   "01B ",
	}, {
		name:   "integer precision 2 format",
		line:   testline(),
		bytes:  1,
		format: "%-03.2d",
		want:   "01B",
	}, {
		name:   "integer precision 2 format",
		line:   testline(),
		bytes:  1,
		format: "%-04.2d",
		want:   "01B ",
	}, {
		name:   "integer precision 2 format",
		line:   testline(),
		bytes:  1,
		format: "%+.2d",
		want:   "+01B",
	}, {
		name:   "integer precision 2 format",
		line:   testline(),
		bytes:  1,
		format: "%#.2d",
		want:   "01B",
	}, {
		name:   "integer precision 2 format",
		line:   testline(),
		bytes:  1024,
		format: "%.2d",
		want:   "01K",
	}, {
		name:   "integer precision 2 format",
		line:   testline(),
		bytes:  1024,
		format: "% .2d",
		want:   "01 K",
		bench:  true,
	}, {
		name:   "integer precision 2 format",
		line:   testline(),
		bytes:  1024,
		format: "% 2.2d",
		want:   "01  K",
	}, {
		name:   "integer precision 2 format",
		line:   testline(),
		bytes:  1024,
		format: "%3.2d",
		want:   "01K",
	}, {
		name:   "integer precision 2 format",
		line:   testline(),
		bytes:  1024,
		format: "%4.2d",
		want:   " 01K",
	}, {
		name:   "integer precision 2 format",
		line:   testline(),
		bytes:  1024,
		format: "%03.2d",
		want:   "01K",
	}, {
		name:   "integer precision 2 format",
		line:   testline(),
		bytes:  1024,
		format: "%04.2d",
		want:   "001K",
	}, {
		name:   "integer precision 2 format",
		line:   testline(),
		bytes:  1024,
		format: "%-3.2d",
		want:   "01K",
	}, {
		name:   "integer precision 2 format",
		line:   testline(),
		bytes:  1024,
		format: "%-4.2d",
		want:   "01K ",
	}, {
		name:   "integer precision 2 format",
		line:   testline(),
		bytes:  1024,
		format: "%-03.2d",
		want:   "01K",
	}, {
		name:   "integer precision 2 format",
		line:   testline(),
		bytes:  1024,
		format: "%-04.2d",
		want:   "01K ",
	}, {
		name:   "integer precision 2 format",
		line:   testline(),
		bytes:  1024,
		format: "%+.2d",
		want:   "+01K",
	}, {
		name:   "integer precision 2 format",
		line:   testline(),
		bytes:  1024,
		format: "%#.2d",
		want:   "01K",
	}, {
		name:   "integer precision 2 format",
		line:   testline(),
		bytes:  1128,
		format: "%.2d",
		want:   "01K",
	}, {
		name:   "integer precision 2 format",
		line:   testline(),
		bytes:  1128,
		format: "% .2d",
		want:   "01 K",
		bench:  true,
	}, {
		name:   "integer precision 2 format",
		line:   testline(),
		bytes:  1128,
		format: "% 2.2d",
		want:   "01  K",
	}, {
		name:   "integer precision 2 format",
		line:   testline(),
		bytes:  1128,
		format: "%3.2d",
		want:   "01K",
	}, {
		name:   "integer precision 2 format",
		line:   testline(),
		bytes:  1128,
		format: "%4.2d",
		want:   " 01K",
	}, {
		name:   "integer precision 2 format",
		line:   testline(),
		bytes:  1128,
		format: "%02.2d",
		want:   "01K",
	}, {
		name:   "integer precision 2 format",
		line:   testline(),
		bytes:  1128,
		format: "%04.2d",
		want:   "001K",
	}, {
		name:   "integer precision 2 format",
		line:   testline(),
		bytes:  1128,
		format: "%-2.2d",
		want:   "01K",
	}, {
		name:   "integer precision 2 format",
		line:   testline(),
		bytes:  1128,
		format: "%-4.2d",
		want:   "01K ",
	}, {
		name:   "integer precision 2 format",
		line:   testline(),
		bytes:  1128,
		format: "%-02.2d",
		want:   "01K",
	}, {
		name:   "integer precision 2 format",
		line:   testline(),
		bytes:  1128,
		format: "%-04.2d",
		want:   "01K ",
	}, {
		name:   "integer precision 2 format",
		line:   testline(),
		bytes:  1128,
		format: "%+.2d",
		want:   "+01K",
	}, {
		name:   "integer precision 2 format",
		line:   testline(),
		bytes:  1128,
		format: "%#.2d",
		want:   "01K",
	}, {
		name:   "name",
		line:   testline(),
		bytes:  1124,
		format: "%d",
		names:  []string{"B", "Kilobyte"},
		want:   "1Kilobyte",
	}, {
		name:   "name",
		line:   testline(),
		bytes:  1124,
		format: "% d",
		names:  []string{"B", "Kilobyte"},
		want:   "1 Kilobyte",
		bench:  true,
	}, {
		name:   "name",
		line:   testline(),
		bytes:  1124,
		format: "% 2d",
		names:  []string{"B", "Kilobyte"},
		want:   "1  Kilobyte",
	},
}

func TestBytesFormat(t *testing.T) {
	for _, tt := range bytesFormatTestc {
		tt := tt
		t.Run(tt.line+"/"+tt.name+" "+tt.format+" "+strconv.FormatUint(tt.bytes, 10), func(t *testing.T) {
			t.Parallel()

			got := fmt.Sprintf(tt.format, bytefmt.New(tt.bytes, tt.names...))
			if got != tt.want {
				t.Errorf("\nwant kilobytes: %#v\n got kilobytes: %#v\ntest: %s", tt.want, got, tt.line)
			}
		})
	}
}

func BenchmarkBytesFormat(b *testing.B) {
	b.ReportAllocs()

	for _, tt := range bytesFormatTestc {
		if !tt.bench {
			continue
		}

		b.Run(tt.line+"/"+tt.name+" "+tt.format+" "+strconv.FormatUint(tt.bytes, 10), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = fmt.Sprintf(tt.format, bytefmt.New(tt.bytes, tt.names...))
			}
		})
	}
}

func testline() string {
	_, file, line, ok := runtime.Caller(1)
	if ok {
		return fmt.Sprintf("%s:%d", filepath.Base(file), line)
	}
	return "it was not possible to recover file and line number information about function invocations"
}
