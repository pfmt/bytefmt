// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bytefmt

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"testing"
)

func testLine() int {
	_, _, line, _ := runtime.Caller(1)
	return line
}

var NamesTestCases = []struct {
	line      int
	names     []string
	expected  []string
	benchmark bool
}{
	{
		line:      testLine(),
		names:     []string{"B", "K", "M", "G", "T", "P", "E"},
		expected:  []string{"B", "K", "M", "G", "T", "P", "E"},
		benchmark: true,
	},
	{
		line:     testLine(),
		names:    []string{},
		expected: []string{"B", "K", "M", "G", "T", "P", "E"},
	},
	{
		line:     testLine(),
		names:    []string{"B"},
		expected: []string{"B", "K", "M", "G", "T", "P", "E"},
	},
	{
		line:     testLine(),
		names:    []string{"B", "K"},
		expected: []string{"B", "K", "M", "G", "T", "P", "E"},
	},
	{
		line:     testLine(),
		names:    []string{"B", "Kilobyte"},
		expected: []string{"B", "Kilobyte", "M", "G", "T", "P", "E"},
	},
	{
		line:     testLine(),
		names:    []string{"B", "K", "M", "G", "T", "P", "E", "Zettabyte"},
		expected: []string{"B", "K", "M", "G", "T", "P", "E", "Zettabyte"},
	},
}

func TestNamesInitialize(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	for _, tc := range NamesTestCases {
		tc := tc
		t.Run(strconv.Itoa(tc.line), func(t *testing.T) {
			t.Parallel()
			linkToExample := fmt.Sprintf("%s:%d", testFile, tc.line)
			b := New(0, tc.names...)
			n := b.Names()
			if strings.Join(n, "") != strings.Join(tc.expected, "") {
				t.Errorf("unexpected name %#v, expected %#v %s", n, tc.expected, linkToExample)
			}
		})
	}
}

func TestNamesUpdate(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	for _, tc := range NamesTestCases {
		tc := tc
		t.Run(strconv.Itoa(tc.line), func(t *testing.T) {
			t.Parallel()
			linkToExample := fmt.Sprintf("%s:%d", testFile, tc.line)
			b := New(0)
			n := b.Names(tc.names...)
			if strings.Join(n, "") != strings.Join(tc.expected, "") {
				t.Errorf("unexpected name %#v, expected %#v %s", n, tc.expected, linkToExample)
			}
		})
	}
}

func BenchmarkNamesInitialize(b *testing.B) {
	b.ReportAllocs()
	for _, tc := range NamesTestCases {
		if !tc.benchmark {
			continue
		}
		b.Run(strconv.Itoa(tc.line), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = New(0, tc.names...)
			}
		})
	}
}

func BenchmarkNamesUpdate(b *testing.B) {
	b.ReportAllocs()
	for _, tc := range NamesTestCases {
		if !tc.benchmark {
			continue
		}
		b.Run(strconv.Itoa(tc.line), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b := New(0)
				_ = b.Names(tc.names...)
			}
		})
	}
}

var BytesKilobytesTestCases = []struct {
	name      string
	line      int
	bytes     uint64
	expected  float64
	benchmark bool
}{
	{
		name:     "exactly one kilobyte",
		line:     testLine(),
		bytes:    1024,
		expected: 1,
	},
	{
		name:     "less than one kilobyte",
		line:     testLine(),
		bytes:    42,
		expected: 0.041015625,
	},
	{
		name:      "more than one kilobyte",
		line:      testLine(),
		bytes:     100500,
		expected:  98.14453125,
		benchmark: true,
	},
}

func TestBytesKilobytes(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	for _, tc := range BytesKilobytesTestCases {
		tc := tc
		t.Run(tc.name+" "+strconv.FormatUint(tc.bytes, 10)+" "+strconv.Itoa(tc.line), func(t *testing.T) {
			t.Parallel()
			linkToExample := fmt.Sprintf("%s:%d", testFile, tc.line)
			f := Bytes{Value: tc.bytes}.kilobytes()
			if f != tc.expected {
				t.Errorf("unexpected kilobytes %v, expected %v %s", f, tc.expected, linkToExample)
			}
		})
	}
}

func BenchmarkBytesKilobytes(b *testing.B) {
	b.ReportAllocs()
	for _, tc := range BytesKilobytesTestCases {
		if !tc.benchmark {
			continue
		}
		b.Run(tc.name+" "+strconv.FormatUint(tc.bytes, 10)+" "+strconv.Itoa(tc.line), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = Bytes{Value: tc.bytes}.kilobytes()
			}
		})
	}
}

var BytesMegabytesTestCases = []struct {
	name      string
	line      int
	bytes     uint64
	expected  float64
	benchmark bool
}{
	{
		name:     "exactly one megabyte",
		line:     testLine(),
		bytes:    1048576,
		expected: 1,
	},
	{
		name:     "less than one megabyte",
		line:     testLine(),
		bytes:    42,
		expected: 4.00543212890625e-05,
	},
	{
		name:      "more than one megabyte",
		line:      testLine(),
		bytes:     10050000,
		expected:  9.584426879882812,
		benchmark: true,
	},
}

func TestBytesMegabytes(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	for _, tc := range BytesMegabytesTestCases {
		tc := tc
		t.Run(tc.name+" "+strconv.FormatUint(tc.bytes, 10)+" "+strconv.Itoa(tc.line), func(t *testing.T) {
			t.Parallel()
			linkToExample := fmt.Sprintf("%s:%d", testFile, tc.line)
			f := Bytes{Value: tc.bytes}.megabytes()
			if f != tc.expected {
				t.Errorf("unexpected megabytes %v, expected %v %s", f, tc.expected, linkToExample)
			}
		})
	}
}

func BenchmarkBytesMegabytes(b *testing.B) {
	b.ReportAllocs()
	for _, tc := range BytesMegabytesTestCases {
		if !tc.benchmark {
			continue
		}
		b.Run(tc.name+" "+strconv.FormatUint(tc.bytes, 10)+" "+strconv.Itoa(tc.line), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = Bytes{Value: tc.bytes}.megabytes()
			}
		})
	}
}

var BytesGigabytesTestCases = []struct {
	name      string
	line      int
	bytes     uint64
	expected  float64
	benchmark bool
}{
	{
		name:     "exactly one gigabyte",
		line:     testLine(),
		bytes:    1073741824,
		expected: 1,
	},
	{
		name:     "less than one gigabyte",
		line:     testLine(),
		bytes:    42,
		expected: 3.91155481338501e-08,
	},
	{
		name:      "more than one gigabyte",
		line:      testLine(),
		bytes:     10050000000,
		expected:  9.359791874885559,
		benchmark: true,
	},
}

func TestBytesGigabytes(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	for _, tc := range BytesGigabytesTestCases {
		tc := tc
		t.Run(tc.name+" "+strconv.FormatUint(tc.bytes, 10)+" "+strconv.Itoa(tc.line), func(t *testing.T) {
			t.Parallel()
			linkToExample := fmt.Sprintf("%s:%d", testFile, tc.line)
			f := Bytes{Value: tc.bytes}.gigabytes()
			if f != tc.expected {
				t.Errorf("unexpected gigabytes %v, expected %v %s", f, tc.expected, linkToExample)
			}
		})
	}
}

func BenchmarkBytesGigabytes(b *testing.B) {
	b.ReportAllocs()
	for _, tc := range BytesGigabytesTestCases {
		if !tc.benchmark {
			continue
		}
		b.Run(tc.name+" "+strconv.FormatUint(tc.bytes, 10)+" "+strconv.Itoa(tc.line), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = Bytes{Value: tc.bytes}.gigabytes()
			}
		})
	}
}

var BytesTerabytesTestCases = []struct {
	name      string
	line      int
	bytes     uint64
	expected  float64
	benchmark bool
}{
	{
		name:     "exactly one terabyte",
		line:     testLine(),
		bytes:    1099511627776,
		expected: 1,
	},
	{
		name:     "less than one terabyte",
		line:     testLine(),
		bytes:    42,
		expected: 3.8198777474462986e-11,
	},
	{
		name:      "more than one terabyte",
		line:      testLine(),
		bytes:     10050000000000,
		expected:  9.140421752817929,
		benchmark: true,
	},
}

func TestBytesTerabytes(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	for _, tc := range BytesTerabytesTestCases {
		tc := tc
		t.Run(tc.name+" "+strconv.FormatUint(tc.bytes, 10)+" "+strconv.Itoa(tc.line), func(t *testing.T) {
			t.Parallel()
			linkToExample := fmt.Sprintf("%s:%d", testFile, tc.line)
			f := Bytes{Value: tc.bytes}.terabytes()
			if f != tc.expected {
				t.Errorf("unexpected terabytes %v, expected %v %s", f, tc.expected, linkToExample)
			}
		})
	}
}

func BenchmarkBytesTerabytes(b *testing.B) {
	b.ReportAllocs()
	for _, tc := range BytesTerabytesTestCases {
		if !tc.benchmark {
			continue
		}
		b.Run(tc.name+" "+strconv.FormatUint(tc.bytes, 10)+" "+strconv.Itoa(tc.line), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = Bytes{Value: tc.bytes}.terabytes()
			}
		})
	}
}

var BytesPetabytesTestCases = []struct {
	name      string
	line      int
	bytes     uint64
	expected  float64
	benchmark bool
}{
	{
		name:     "exactly one petabyte",
		line:     testLine(),
		bytes:    1125899906842624,
		expected: 1,
	},
	{
		name:     "less than one petabyte",
		line:     testLine(),
		bytes:    42,
		expected: 3.730349362740526e-14,
	},
	{
		name:      "more than one petabyte",
		line:      testLine(),
		bytes:     10050000000000000,
		expected:  8.926193117986259,
		benchmark: true,
	},
}

func TestBytesPetabytes(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	for _, tc := range BytesPetabytesTestCases {
		tc := tc
		t.Run(tc.name+" "+strconv.FormatUint(tc.bytes, 10)+" "+strconv.Itoa(tc.line), func(t *testing.T) {
			t.Parallel()
			linkToExample := fmt.Sprintf("%s:%d", testFile, tc.line)
			f := Bytes{Value: tc.bytes}.petabytes()
			if f != tc.expected {
				t.Errorf("unexpected petabytes %v, expected %v %s", f, tc.expected, linkToExample)
			}
		})
	}
}

func BenchmarkBytesPetabytes(b *testing.B) {
	b.ReportAllocs()
	for _, tc := range BytesPetabytesTestCases {
		if !tc.benchmark {
			continue
		}
		b.Run(tc.name+" "+strconv.FormatUint(tc.bytes, 10)+" "+strconv.Itoa(tc.line), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = Bytes{Value: tc.bytes}.petabytes()
			}
		})
	}
}

var BytesExabytesTestCases = []struct {
	name      string
	line      int
	bytes     uint64
	expected  float64
	benchmark bool
}{
	{
		name:     "exactly one exabyte",
		line:     testLine(),
		bytes:    1152921504606846976,
		expected: 1,
	},
	{
		name:     "less than one exabyte",
		line:     testLine(),
		bytes:    42,
		expected: 3.642919299551295e-17,
	},
	{
		name:      "more than one exabyte",
		line:      testLine(),
		bytes:     10050000000000000000,
		expected:  8.716985466783456,
		benchmark: true,
	},
}

func TestBytesExabytes(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	for _, tc := range BytesExabytesTestCases {
		tc := tc
		t.Run(tc.name+" "+strconv.FormatUint(tc.bytes, 10)+" "+strconv.Itoa(tc.line), func(t *testing.T) {
			t.Parallel()
			linkToExample := fmt.Sprintf("%s:%d", testFile, tc.line)
			f := Bytes{Value: tc.bytes}.exabytes()
			if f != tc.expected {
				t.Errorf("unexpected exabytes %v, expected %v %s", f, tc.expected, linkToExample)
			}
		})
	}
}

func BenchmarkBytesExabytes(b *testing.B) {
	b.ReportAllocs()
	for _, tc := range BytesExabytesTestCases {
		if !tc.benchmark {
			continue
		}
		b.Run(tc.name+" "+strconv.FormatUint(tc.bytes, 10)+" "+strconv.Itoa(tc.line), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = Bytes{Value: tc.bytes}.exabytes()
			}
		})
	}
}

var BytesStringTestCases = []struct {
	name      string
	line      int
	bytes     uint64
	expected  string
	benchmark bool
}{
	{
		name:     "zero byte",
		line:     testLine(),
		bytes:    0,
		expected: "0B",
	},
	{
		name:     "less than one kilobyte",
		line:     testLine(),
		bytes:    1,
		expected: "1B",
	},
	{
		name:     "exactly one kilobyte",
		line:     testLine(),
		bytes:    1024,
		expected: "1K",
	},
	{
		name:      "more than one kilobyte",
		line:      testLine(),
		bytes:     1128,
		expected:  "1.1015625K",
		benchmark: true,
	},
}

func TestBytesString(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	for _, tc := range BytesStringTestCases {
		tc := tc
		t.Run(tc.name+" "+strconv.FormatUint(tc.bytes, 10)+" "+strconv.Itoa(tc.line), func(t *testing.T) {
			t.Parallel()
			linkToExample := fmt.Sprintf("%s:%d", testFile, tc.line)
			s := New(tc.bytes).String()
			if s != tc.expected {
				t.Errorf("unexpected string %#v, expected %#v %s",
					s, tc.expected, linkToExample)
			}
		})
	}
}

func BenchmarkBytesString(b *testing.B) {
	b.ReportAllocs()
	for _, tc := range BytesStringTestCases {
		if !tc.benchmark {
			continue
		}
		b.Run(tc.name+" "+strconv.FormatUint(tc.bytes, 10)+" "+strconv.Itoa(tc.line), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = New(tc.bytes).String()
			}
		})
	}
}

var BytesFormatTestCases = []struct {
	name      string
	line      int
	bytes     uint64
	format    string
	names     []string
	expected  string
	benchmark bool
}{
	{
		name:     "general format",
		line:     testLine(),
		bytes:    0,
		format:   "%v",
		expected: "0B",
	},
	{
		name:      "general format",
		line:      testLine(),
		bytes:     0,
		format:    "% v",
		expected:  "0 B",
		benchmark: true,
	},
	{
		name:     "general format",
		line:     testLine(),
		bytes:    0,
		format:   "% 2v",
		expected: "0  B",
	},
	{
		name:     "general format",
		line:     testLine(),
		bytes:    0,
		format:   "%3v",
		expected: " 0B",
	},
	{
		name:     "general format",
		line:     testLine(),
		bytes:    0,
		format:   "%03v",
		expected: "00B",
	},
	{
		name:     "general format",
		line:     testLine(),
		bytes:    0,
		format:   "%-3v",
		expected: "0B ",
	},
	{
		name:     "general format",
		line:     testLine(),
		bytes:    0,
		format:   "%-03v",
		expected: "0B ",
	},
	{
		name:     "general format",
		line:     testLine(),
		bytes:    0,
		format:   "%+v",
		expected: "0B",
	},
	{
		name:     "general format",
		line:     testLine(),
		bytes:    0,
		format:   "%#v",
		expected: "0B",
	},

	{
		name:     "general format",
		line:     testLine(),
		bytes:    1,
		format:   "%v",
		expected: "1B",
	},
	{
		name:      "general format",
		line:      testLine(),
		bytes:     1,
		format:    "% v",
		expected:  "1 B",
		benchmark: true,
	},
	{
		name:     "general format",
		line:     testLine(),
		bytes:    1,
		format:   "% 2v",
		expected: "1  B",
	},
	{
		name:     "general format",
		line:     testLine(),
		bytes:    1,
		format:   "%3v",
		expected: " 1B",
	},
	{
		name:     "general format",
		line:     testLine(),
		bytes:    1,
		format:   "%03v",
		expected: "01B",
	},
	{
		name:     "general format",
		line:     testLine(),
		bytes:    1,
		format:   "%-3v",
		expected: "1B ",
	},
	{
		name:     "general format",
		line:     testLine(),
		bytes:    1,
		format:   "%-03v",
		expected: "1B ",
	},
	{
		name:     "general format",
		line:     testLine(),
		bytes:    1,
		format:   "%+v",
		expected: "1B",
	},
	{
		name:     "general format",
		line:     testLine(),
		bytes:    1,
		format:   "%#v",
		expected: "1B",
	},

	{
		name:     "general format",
		line:     testLine(),
		bytes:    1024,
		format:   "%v",
		expected: "1K",
	},
	{
		name:      "general format",
		line:      testLine(),
		bytes:     1024,
		format:    "% v",
		expected:  "1 K",
		benchmark: true,
	},
	{
		name:     "general format",
		line:     testLine(),
		bytes:    1024,
		format:   "% 2v",
		expected: "1  K",
	},
	{
		name:     "general format",
		line:     testLine(),
		bytes:    1024,
		format:   "%3v",
		expected: " 1K",
	},
	{
		name:     "general format",
		line:     testLine(),
		bytes:    1024,
		format:   "%03v",
		expected: "01K",
	},
	{
		name:     "general format",
		line:     testLine(),
		bytes:    1024,
		format:   "%-3v",
		expected: "1K ",
	},
	{
		name:     "general format",
		line:     testLine(),
		bytes:    1024,
		format:   "%-03v",
		expected: "1K ",
	},
	{
		name:     "general format",
		line:     testLine(),
		bytes:    1024,
		format:   "%+v",
		expected: "1K",
	},
	{
		name:     "general format",
		line:     testLine(),
		bytes:    1024,
		format:   "%#v",
		expected: "1K",
	},

	{
		name:     "general format",
		line:     testLine(),
		bytes:    1128,
		format:   "%v",
		expected: "1.1015625K",
	},
	{
		name:      "general format",
		line:      testLine(),
		bytes:     1128,
		format:    "% v",
		expected:  "1.1015625 K",
		benchmark: true,
	},
	{
		name:     "general format",
		line:     testLine(),
		bytes:    1128,
		format:   "% 2v",
		expected: "1.1015625  K",
	},
	{
		name:     "general format",
		line:     testLine(),
		bytes:    1128,
		format:   "%11v",
		expected: " 1.1015625K",
	},
	{
		name:     "general format",
		line:     testLine(),
		bytes:    1128,
		format:   "%011v",
		expected: "01.1015625K",
	},
	{
		name:     "general format",
		line:     testLine(),
		bytes:    1128,
		format:   "%-11v",
		expected: "1.1015625K ",
	},
	{
		name:     "general format",
		line:     testLine(),
		bytes:    1128,
		format:   "%-011v",
		expected: "1.1015625K ",
	},
	{
		name:     "general format",
		line:     testLine(),
		bytes:    1128,
		format:   "%+v",
		expected: "1.1015625K",
	},
	{
		name:     "general format",
		line:     testLine(),
		bytes:    1128,
		format:   "%#v",
		expected: "1.1015625K",
	},
	{
		name:     "string format",
		line:     testLine(),
		bytes:    0,
		format:   "%s",
		expected: "0B",
	},
	{
		name:      "string format",
		line:      testLine(),
		bytes:     0,
		format:    "% s",
		expected:  "0 B",
		benchmark: true,
	},
	{
		name:     "string format",
		line:     testLine(),
		bytes:    0,
		format:   "% 2s",
		expected: "0  B",
	},
	{
		name:     "string format",
		line:     testLine(),
		bytes:    0,
		format:   "%3s",
		expected: " 0B",
	},
	{
		name:     "string format",
		line:     testLine(),
		bytes:    0,
		format:   "%03s",
		expected: "00B",
	},
	{
		name:     "string format",
		line:     testLine(),
		bytes:    0,
		format:   "%-3s",
		expected: "0B ",
	},
	{
		name:     "string format",
		line:     testLine(),
		bytes:    0,
		format:   "%-03s",
		expected: "0B ",
	},
	{
		name:     "string format",
		line:     testLine(),
		bytes:    0,
		format:   "%+s",
		expected: "0B",
	},
	{
		name:     "string format",
		line:     testLine(),
		bytes:    0,
		format:   "%#s",
		expected: "0B",
	},

	{
		name:     "string format",
		line:     testLine(),
		bytes:    1,
		format:   "%s",
		expected: "1B",
	},
	{
		name:      "string format",
		line:      testLine(),
		bytes:     1,
		format:    "% s",
		expected:  "1 B",
		benchmark: true,
	},
	{
		name:     "string format",
		line:     testLine(),
		bytes:    1,
		format:   "% 2s",
		expected: "1  B",
	},
	{
		name:     "string format",
		line:     testLine(),
		bytes:    1,
		format:   "%3s",
		expected: " 1B",
	},
	{
		name:     "string format",
		line:     testLine(),
		bytes:    1,
		format:   "%03s",
		expected: "01B",
	},
	{
		name:     "string format",
		line:     testLine(),
		bytes:    1,
		format:   "%-3s",
		expected: "1B ",
	},
	{
		name:     "string format",
		line:     testLine(),
		bytes:    1,
		format:   "%-03s",
		expected: "1B ",
	},
	{
		name:     "string format",
		line:     testLine(),
		bytes:    1,
		format:   "%+s",
		expected: "1B",
	},
	{
		name:     "string format",
		line:     testLine(),
		bytes:    1,
		format:   "%#s",
		expected: "1B",
	},

	{
		name:     "string format",
		line:     testLine(),
		bytes:    1024,
		format:   "%s",
		expected: "1K",
	},
	{
		name:      "string format",
		line:      testLine(),
		bytes:     1024,
		format:    "% s",
		expected:  "1 K",
		benchmark: true,
	},
	{
		name:     "string format",
		line:     testLine(),
		bytes:    1024,
		format:   "% 2s",
		expected: "1  K",
	},
	{
		name:     "string format",
		line:     testLine(),
		bytes:    1024,
		format:   "%3s",
		expected: " 1K",
	},
	{
		name:     "string format",
		line:     testLine(),
		bytes:    1024,
		format:   "%03s",
		expected: "01K",
	},
	{
		name:     "string format",
		line:     testLine(),
		bytes:    1024,
		format:   "%-3s",
		expected: "1K ",
	},
	{
		name:     "string format",
		line:     testLine(),
		bytes:    1024,
		format:   "%-03s",
		expected: "1K ",
	},
	{
		name:     "string format",
		line:     testLine(),
		bytes:    1024,
		format:   "%+s",
		expected: "1K",
	},
	{
		name:     "string format",
		line:     testLine(),
		bytes:    1024,
		format:   "%#s",
		expected: "1K",
	},

	{
		name:     "string format",
		line:     testLine(),
		bytes:    1128,
		format:   "%s",
		expected: "1.1015625K",
	},
	{
		name:      "string format",
		line:      testLine(),
		bytes:     1128,
		format:    "% s",
		expected:  "1.1015625 K",
		benchmark: true,
	},
	{
		name:     "string format",
		line:     testLine(),
		bytes:    1128,
		format:   "% 2s",
		expected: "1.1015625  K",
	},
	{
		name:     "string format",
		line:     testLine(),
		bytes:    1128,
		format:   "%11s",
		expected: " 1.1015625K",
	},
	{
		name:     "string format",
		line:     testLine(),
		bytes:    1128,
		format:   "%011s",
		expected: "01.1015625K",
	},
	{
		name:     "string format",
		line:     testLine(),
		bytes:    1128,
		format:   "%-11s",
		expected: "1.1015625K ",
	},
	{
		name:     "string format",
		line:     testLine(),
		bytes:    1128,
		format:   "%-011s",
		expected: "1.1015625K ",
	},
	{
		name:     "string format",
		line:     testLine(),
		bytes:    1128,
		format:   "%+s",
		expected: "1.1015625K",
	},
	{
		name:     "string format",
		line:     testLine(),
		bytes:    1128,
		format:   "%#s",
		expected: "1.1015625K",
	},
	{
		name:     "double-quoted string format",
		line:     testLine(),
		bytes:    0,
		format:   "%q",
		expected: `"0B"`,
	},
	{
		name:      "double-quoted string format",
		line:      testLine(),
		bytes:     0,
		format:    "% q",
		expected:  `"0 B"`,
		benchmark: true,
	},
	{
		name:     "double-quoted string format",
		line:     testLine(),
		bytes:    0,
		format:   "% 2q",
		expected: `"0  B"`,
	},
	{
		name:     "double-quoted string format",
		line:     testLine(),
		bytes:    0,
		format:   "%3q",
		expected: `" 0B"`,
	},
	{
		name:     "double-quoted string format",
		line:     testLine(),
		bytes:    0,
		format:   "%03q",
		expected: `"00B"`,
	},
	{
		name:     "double-quoted string format",
		line:     testLine(),
		bytes:    0,
		format:   "%-3q",
		expected: `"0B "`,
	},
	{
		name:     "double-quoted string format",
		line:     testLine(),
		bytes:    0,
		format:   "%-03q",
		expected: `"0B "`,
	},
	{
		name:     "double-quoted string format",
		line:     testLine(),
		bytes:    0,
		format:   "%+q",
		expected: `"0B"`,
	},
	{
		name:     "double-quoted string format",
		line:     testLine(),
		bytes:    0,
		format:   "%#q",
		expected: "`0B`",
	},

	{
		name:     "double-quoted string format",
		line:     testLine(),
		bytes:    1,
		format:   "%q",
		expected: `"1B"`,
	},
	{
		name:      "double-quoted string format",
		line:      testLine(),
		bytes:     1,
		format:    "% q",
		expected:  `"1 B"`,
		benchmark: true,
	},
	{
		name:     "double-quoted string format",
		line:     testLine(),
		bytes:    1,
		format:   "% 2q",
		expected: `"1  B"`,
	},
	{
		name:     "double-quoted string format",
		line:     testLine(),
		bytes:    1,
		format:   "%3q",
		expected: `" 1B"`,
	},
	{
		name:     "double-quoted string format",
		line:     testLine(),
		bytes:    1,
		format:   "%03q",
		expected: `"01B"`,
	},
	{
		name:     "double-quoted string format",
		line:     testLine(),
		bytes:    1,
		format:   "%-3q",
		expected: `"1B "`,
	},
	{
		name:     "double-quoted string format",
		line:     testLine(),
		bytes:    1,
		format:   "%-03q",
		expected: `"1B "`,
	},
	{
		name:     "double-quoted string format",
		line:     testLine(),
		bytes:    1,
		format:   "%+q",
		expected: `"1B"`,
	},
	{
		name:     "double-quoted string format",
		line:     testLine(),
		bytes:    1,
		format:   "%#q",
		expected: "`1B`",
	},

	{
		name:     "double-quoted string format",
		line:     testLine(),
		bytes:    1024,
		format:   "%q",
		expected: `"1K"`,
	},
	{
		name:      "double-quoted string format",
		line:      testLine(),
		bytes:     1024,
		format:    "% q",
		expected:  `"1 K"`,
		benchmark: true,
	},
	{
		name:     "double-quoted string format",
		line:     testLine(),
		bytes:    1024,
		format:   "% 2q",
		expected: `"1  K"`,
	},
	{
		name:     "double-quoted string format",
		line:     testLine(),
		bytes:    1024,
		format:   "%3q",
		expected: `" 1K"`,
	},
	{
		name:     "double-quoted string format",
		line:     testLine(),
		bytes:    1024,
		format:   "%03q",
		expected: `"01K"`,
	},
	{
		name:     "double-quoted string format",
		line:     testLine(),
		bytes:    1024,
		format:   "%-3q",
		expected: `"1K "`,
	},
	{
		name:     "double-quoted string format",
		line:     testLine(),
		bytes:    1024,
		format:   "%-03q",
		expected: `"1K "`,
	},
	{
		name:     "double-quoted string format",
		line:     testLine(),
		bytes:    1024,
		format:   "%+q",
		expected: `"1K"`,
	},
	{
		name:     "double-quoted string format",
		line:     testLine(),
		bytes:    1024,
		format:   "%#q",
		expected: "`1K`",
	},

	{
		name:     "double-quoted string format",
		line:     testLine(),
		bytes:    1128,
		format:   "%q",
		expected: `"1.1015625K"`,
	},
	{
		name:      "double-quoted string format",
		line:      testLine(),
		bytes:     1128,
		format:    "% q",
		expected:  `"1.1015625 K"`,
		benchmark: true,
	},
	{
		name:     "double-quoted string format",
		line:     testLine(),
		bytes:    1128,
		format:   "% 2q",
		expected: `"1.1015625  K"`,
	},
	{
		name:     "double-quoted string format",
		line:     testLine(),
		bytes:    1128,
		format:   "%11q",
		expected: `" 1.1015625K"`,
	},
	{
		name:     "double-quoted string format",
		line:     testLine(),
		bytes:    1128,
		format:   "%011q",
		expected: `"01.1015625K"`,
	},
	{
		name:     "double-quoted string format",
		line:     testLine(),
		bytes:    1128,
		format:   "%-11q",
		expected: `"1.1015625K "`,
	},
	{
		name:     "double-quoted string format",
		line:     testLine(),
		bytes:    1128,
		format:   "%-011q",
		expected: `"1.1015625K "`,
	},
	{
		name:     "double-quoted string format",
		line:     testLine(),
		bytes:    1128,
		format:   "%+q",
		expected: `"1.1015625K"`,
	},
	{
		name:     "double-quoted string format",
		line:     testLine(),
		bytes:    1128,
		format:   "%#q",
		expected: "`1.1015625K`",
	},

	{
		name:     "float format",
		line:     testLine(),
		bytes:    0,
		format:   "%f",
		expected: "0.000000B",
	},
	{
		name:      "float format",
		line:      testLine(),
		bytes:     0,
		format:    "% f",
		expected:  "0.000000 B",
		benchmark: true,
	},
	{
		name:     "float format",
		line:     testLine(),
		bytes:    0,
		format:   "% 2f",
		expected: "0.000000  B",
	},
	{
		name:     "float format",
		line:     testLine(),
		bytes:    0,
		format:   "%10f",
		expected: " 0.000000B",
	},
	{
		name:     "float format",
		line:     testLine(),
		bytes:    0,
		format:   "%010f",
		expected: "00.000000B",
	},
	{
		name:     "float format",
		line:     testLine(),
		bytes:    0,
		format:   "%-10f",
		expected: "0.000000B ",
	},
	{
		name:     "float format",
		line:     testLine(),
		bytes:    0,
		format:   "%-010f",
		expected: "0.000000B ",
	},
	{
		name:     "float format",
		line:     testLine(),
		bytes:    0,
		format:   "%+f",
		expected: "+0.000000B",
	},
	{
		name:     "float format",
		line:     testLine(),
		bytes:    0,
		format:   "%#f",
		expected: "0.000000B",
	},
	{
		name:      "float format",
		line:      testLine(),
		bytes:     1,
		format:    "%f",
		expected:  "1.000000B",
		benchmark: true,
	},
	{
		name:     "float format",
		line:     testLine(),
		bytes:    1,
		format:   "% 2f",
		expected: "1.000000  B",
	},
	{
		name:     "float format",
		line:     testLine(),
		bytes:    1,
		format:   "%10f",
		expected: " 1.000000B",
	},
	{
		name:     "float format",
		line:     testLine(),
		bytes:    1,
		format:   "%010f",
		expected: "01.000000B",
	},
	{
		name:     "float format",
		line:     testLine(),
		bytes:    1,
		format:   "%-10f",
		expected: "1.000000B ",
	},
	{
		name:     "float format",
		line:     testLine(),
		bytes:    1,
		format:   "%-010f",
		expected: "1.000000B ",
	},
	{
		name:     "float format",
		line:     testLine(),
		bytes:    1,
		format:   "%+f",
		expected: "+1.000000B",
	},
	{
		name:     "float format",
		line:     testLine(),
		bytes:    1,
		format:   "%#f",
		expected: "1.000000B",
	},
	{
		name:     "float format",
		line:     testLine(),
		bytes:    1024,
		format:   "%f",
		expected: "1.000000K",
	},
	{
		name:      "float format",
		line:      testLine(),
		bytes:     1024,
		format:    "% f",
		expected:  "1.000000 K",
		benchmark: true,
	},
	{
		name:     "float format",
		line:     testLine(),
		bytes:    1024,
		format:   "% 2f",
		expected: "1.000000  K",
	},
	{
		name:     "float format",
		line:     testLine(),
		bytes:    1024,
		format:   "%10f",
		expected: " 1.000000K",
	},
	{
		name:     "float format",
		line:     testLine(),
		bytes:    1024,
		format:   "%010f",
		expected: "01.000000K",
	},
	{
		name:     "float format",
		line:     testLine(),
		bytes:    1024,
		format:   "%-10f",
		expected: "1.000000K ",
	},
	{
		name:     "float format",
		line:     testLine(),
		bytes:    1024,
		format:   "%-010f",
		expected: "1.000000K ",
	},
	{
		name:     "float format",
		line:     testLine(),
		bytes:    1024,
		format:   "%+f",
		expected: "+1.000000K",
	},
	{
		name:     "float format",
		line:     testLine(),
		bytes:    1024,
		format:   "%#f",
		expected: "1.000000K",
	},
	{
		name:     "float format",
		line:     testLine(),
		bytes:    1128,
		format:   "%f",
		expected: "1.101562K",
	},
	{
		name:      "float format",
		line:      testLine(),
		bytes:     1128,
		format:    "% f",
		expected:  "1.101562 K",
		benchmark: true,
	},
	{
		name:     "float format",
		line:     testLine(),
		bytes:    1128,
		format:   "% 2f",
		expected: "1.101562  K",
	},
	{
		name:     "float format",
		line:     testLine(),
		bytes:    1128,
		format:   "%10f",
		expected: " 1.101562K",
	},
	{
		name:     "float format",
		line:     testLine(),
		bytes:    1128,
		format:   "%010f",
		expected: "01.101562K",
	},
	{
		name:     "float format",
		line:     testLine(),
		bytes:    1128,
		format:   "%-10f",
		expected: "1.101562K ",
	},
	{
		name:     "float format",
		line:     testLine(),
		bytes:    1128,
		format:   "%-010f",
		expected: "1.101562K ",
	},
	{
		name:     "float format",
		line:     testLine(),
		bytes:    1128,
		format:   "%+f",
		expected: "+1.101562K",
	},
	{
		name:     "float format",
		line:     testLine(),
		bytes:    1128,
		format:   "%#f",
		expected: "1.101562K",
	},
	{
		name:     "float precision 1 format",
		line:     testLine(),
		bytes:    0,
		format:   "%.1f",
		expected: "0.0B",
	},
	{
		name:      "float precision 1 format",
		line:      testLine(),
		bytes:     0,
		format:    "% .1f",
		expected:  "0.0 B",
		benchmark: true,
	},
	{
		name:     "float precision 1 format",
		line:     testLine(),
		bytes:    0,
		format:   "% 2.1f",
		expected: "0.0  B",
	},
	{
		name:     "float precision 1 format",
		line:     testLine(),
		bytes:    0,
		format:   "%4.1f",
		expected: "0.0B",
	},
	{
		name:     "float precision 1 format",
		line:     testLine(),
		bytes:    0,
		format:   "%6.1f",
		expected: "  0.0B",
	},
	{
		name:     "float precision 1 format",
		line:     testLine(),
		bytes:    0,
		format:   "%04.1f",
		expected: "0.0B",
	},
	{
		name:     "float precision 1 format",
		line:     testLine(),
		bytes:    0,
		format:   "%06.1f",
		expected: "000.0B",
	},
	{
		name:     "float precision 1 format",
		line:     testLine(),
		bytes:    0,
		format:   "%-5.1f",
		expected: "0.0B ",
	},
	{
		name:     "float precision 1 format",
		line:     testLine(),
		bytes:    0,
		format:   "%-05.1f",
		expected: "0.0B ",
	},
	{
		name:     "float precision 1 format",
		line:     testLine(),
		bytes:    0,
		format:   "%+.1f",
		expected: "+0.0B",
	},
	{
		name:     "float precision 1 format",
		line:     testLine(),
		bytes:    0,
		format:   "%#.1f",
		expected: "0.0B",
	},
	{
		name:     "float precision 1 format",
		line:     testLine(),
		bytes:    1,
		format:   "%.1f",
		expected: "1.0B",
	},
	{
		name:      "float precision 1 format",
		line:      testLine(),
		bytes:     1,
		format:    "% .1f",
		expected:  "1.0 B",
		benchmark: true,
	},
	{
		name:     "float precision 1 format",
		line:     testLine(),
		bytes:    1,
		format:   "% 2.1f",
		expected: "1.0  B",
	},
	{
		name:     "float precision 1 format",
		line:     testLine(),
		bytes:    1,
		format:   "%5.1f",
		expected: " 1.0B",
	},
	{
		name:     "float precision 1 format",
		line:     testLine(),
		bytes:    1,
		format:   "%05.1f",
		expected: "01.0B",
	},
	{
		name:     "float precision 1 format",
		line:     testLine(),
		bytes:    1,
		format:   "%-5.1f",
		expected: "1.0B ",
	},
	{
		name:     "float precision 1 format",
		line:     testLine(),
		bytes:    1,
		format:   "%-05.1f",
		expected: "1.0B ",
	},
	{
		name:     "float precision 1 format",
		line:     testLine(),
		bytes:    1,
		format:   "%+.1f",
		expected: "+1.0B",
	},
	{
		name:     "float precision 1 format",
		line:     testLine(),
		bytes:    1,
		format:   "%#.1f",
		expected: "1.0B",
	},
	{
		name:     "float precision 1 format",
		line:     testLine(),
		bytes:    1024,
		format:   "%.1f",
		expected: "1.0K",
	},
	{
		name:      "float precision 1 format",
		line:      testLine(),
		bytes:     1024,
		format:    "% .1f",
		expected:  "1.0 K",
		benchmark: true,
	},
	{
		name:     "float precision 1 format",
		line:     testLine(),
		bytes:    1024,
		format:   "% 2.1f",
		expected: "1.0  K",
	},
	{
		name:     "float precision 1 format",
		line:     testLine(),
		bytes:    1024,
		format:   "%5.1f",
		expected: " 1.0K",
	},
	{
		name:     "float precision 1 format",
		line:     testLine(),
		bytes:    1024,
		format:   "%05.1f",
		expected: "01.0K",
	},
	{
		name:     "float precision 1 format",
		line:     testLine(),
		bytes:    1024,
		format:   "%-5.1f",
		expected: "1.0K ",
	},
	{
		name:     "float precision 1 format",
		line:     testLine(),
		bytes:    1024,
		format:   "%-05.1f",
		expected: "1.0K ",
	},
	{
		name:     "float precision 1 format",
		line:     testLine(),
		bytes:    1024,
		format:   "%+.1f",
		expected: "+1.0K",
	},
	{
		name:     "float precision 1 format",
		line:     testLine(),
		bytes:    1024,
		format:   "%#.1f",
		expected: "1.0K",
	},
	{
		name:     "float precision 1 format",
		line:     testLine(),
		bytes:    1128,
		format:   "%.1f",
		expected: "1.1K",
	},
	{
		name:      "float precision 1 format",
		line:      testLine(),
		bytes:     1128,
		format:    "% .1f",
		expected:  "1.1 K",
		benchmark: true,
	},
	{
		name:     "float precision 1 format",
		line:     testLine(),
		bytes:    1128,
		format:   "% 2.1f",
		expected: "1.1  K",
	},
	{
		name:     "float precision 1 format",
		line:     testLine(),
		bytes:    1128,
		format:   "%5.1f",
		expected: " 1.1K",
	},
	{
		name:     "float precision 1 format",
		line:     testLine(),
		bytes:    1128,
		format:   "%05.1f",
		expected: "01.1K",
	},
	{
		name:     "float precision 1 format",
		line:     testLine(),
		bytes:    1128,
		format:   "%-5.1f",
		expected: "1.1K ",
	},
	{
		name:     "float precision 1 format",
		line:     testLine(),
		bytes:    1128,
		format:   "%-05.1f",
		expected: "1.1K ",
	},
	{
		name:     "float precision 1 format",
		line:     testLine(),
		bytes:    1128,
		format:   "%+.1f",
		expected: "+1.1K",
	},
	{
		name:     "float precision 1 format",
		line:     testLine(),
		bytes:    1128,
		format:   "%#.1f",
		expected: "1.1K",
	},
	{
		name:     "integer format",
		line:     testLine(),
		bytes:    0,
		format:   "%d",
		expected: "0B",
	},
	{
		name:      "integer format",
		line:      testLine(),
		bytes:     0,
		format:    "% d",
		expected:  "0 B",
		benchmark: true,
	},
	{
		name:     "integer format",
		line:     testLine(),
		bytes:    0,
		format:   "% 2d",
		expected: "0  B",
	},
	{
		name:     "integer format",
		line:     testLine(),
		bytes:    0,
		format:   "%3d",
		expected: " 0B",
	},
	{
		name:     "integer format",
		line:     testLine(),
		bytes:    0,
		format:   "%03d",
		expected: "00B",
	},
	{
		name:     "integer format",
		line:     testLine(),
		bytes:    0,
		format:   "%-3d",
		expected: "0B ",
	},
	{
		name:     "integer format",
		line:     testLine(),
		bytes:    0,
		format:   "%-03d",
		expected: "0B ",
	},
	{
		name:     "integer format",
		line:     testLine(),
		bytes:    0,
		format:   "%+d",
		expected: "+0B",
	},
	{
		name:     "integer format",
		line:     testLine(),
		bytes:    0,
		format:   "%#d",
		expected: "0B",
	},
	{
		name:     "integer format",
		line:     testLine(),
		bytes:    1,
		format:   "%d",
		expected: "1B",
	},
	{
		name:      "integer format",
		line:      testLine(),
		bytes:     1,
		format:    "% d",
		expected:  "1 B",
		benchmark: true,
	},
	{
		name:     "integer format",
		line:     testLine(),
		bytes:    1,
		format:   "% 2d",
		expected: "1  B",
	},
	{
		name:     "integer format",
		line:     testLine(),
		bytes:    1,
		format:   "%3d",
		expected: " 1B",
	},
	{
		name:     "integer format",
		line:     testLine(),
		bytes:    1,
		format:   "%03d",
		expected: "01B",
	},
	{
		name:     "integer format",
		line:     testLine(),
		bytes:    1,
		format:   "%-3d",
		expected: "1B ",
	},
	{
		name:     "integer format",
		line:     testLine(),
		bytes:    1,
		format:   "%-03d",
		expected: "1B ",
	},
	{
		name:     "integer format",
		line:     testLine(),
		bytes:    1,
		format:   "%+d",
		expected: "+1B",
	},
	{
		name:     "integer format",
		line:     testLine(),
		bytes:    1,
		format:   "%#d",
		expected: "1B",
	},
	{
		name:     "integer format",
		line:     testLine(),
		bytes:    1024,
		format:   "%d",
		expected: "1K",
	},
	{
		name:      "integer format",
		line:      testLine(),
		bytes:     1024,
		format:    "% d",
		expected:  "1 K",
		benchmark: true,
	},
	{
		name:     "integer format",
		line:     testLine(),
		bytes:    1024,
		format:   "% 2d",
		expected: "1  K",
	},
	{
		name:     "integer format",
		line:     testLine(),
		bytes:    1024,
		format:   "%3d",
		expected: " 1K",
	},
	{
		name:     "integer format",
		line:     testLine(),
		bytes:    1024,
		format:   "%03d",
		expected: "01K",
	},
	{
		name:     "integer format",
		line:     testLine(),
		bytes:    1024,
		format:   "%-3d",
		expected: "1K ",
	},
	{
		name:     "integer format",
		line:     testLine(),
		bytes:    1024,
		format:   "%-03d",
		expected: "1K ",
	},
	{
		name:     "integer format",
		line:     testLine(),
		bytes:    1024,
		format:   "%+d",
		expected: "+1K",
	},
	{
		name:     "integer format",
		line:     testLine(),
		bytes:    1024,
		format:   "%#d",
		expected: "1K",
	},
	{
		name:     "integer format",
		line:     testLine(),
		bytes:    1128,
		format:   "%d",
		expected: "1K",
	},
	{
		name:      "integer format",
		line:      testLine(),
		bytes:     1128,
		format:    "% d",
		expected:  "1 K",
		benchmark: true,
	},
	{
		name:     "integer format",
		line:     testLine(),
		bytes:    1128,
		format:   "% 2d",
		expected: "1  K",
	},
	{
		name:     "integer format",
		line:     testLine(),
		bytes:    1128,
		format:   "%3d",
		expected: " 1K",
	},
	{
		name:     "integer format",
		line:     testLine(),
		bytes:    1128,
		format:   "%03d",
		expected: "01K",
	},
	{
		name:     "integer format",
		line:     testLine(),
		bytes:    1128,
		format:   "%-3d",
		expected: "1K ",
	},
	{
		name:     "integer format",
		line:     testLine(),
		bytes:    1128,
		format:   "%-03d",
		expected: "1K ",
	},
	{
		name:     "integer format",
		line:     testLine(),
		bytes:    1128,
		format:   "%+d",
		expected: "+1K",
	},
	{
		name:     "integer format",
		line:     testLine(),
		bytes:    1128,
		format:   "%#d",
		expected: "1K",
	},
	{
		name:     "integer precision 2 format",
		line:     testLine(),
		bytes:    0,
		format:   "%.2d",
		expected: "00B",
	},
	{
		name:      "integer precision 2 format",
		line:      testLine(),
		bytes:     0,
		format:    "% .2d",
		expected:  "00 B",
		benchmark: true,
	},
	{
		name:     "integer precision 2 format",
		line:     testLine(),
		bytes:    0,
		format:   "% 2.2d",
		expected: "00  B",
	},
	{
		name:     "integer precision 2 format",
		line:     testLine(),
		bytes:    0,
		format:   "%3.2d",
		expected: "00B",
	},
	{
		name:     "integer precision 2 format",
		line:     testLine(),
		bytes:    0,
		format:   "%4.2d",
		expected: " 00B",
	},
	{
		name:     "integer precision 2 format",
		line:     testLine(),
		bytes:    0,
		format:   "%03.2d",
		expected: "00B",
	},
	{
		name:     "integer precision 2 format",
		line:     testLine(),
		bytes:    0,
		format:   "%-3.2d",
		expected: "00B",
	},
	{
		name:     "integer precision 2 format",
		line:     testLine(),
		bytes:    0,
		format:   "%-4.2d",
		expected: "00B ",
	},
	{
		name:     "integer precision 2 format",
		line:     testLine(),
		bytes:    0,
		format:   "%-03.2d",
		expected: "00B",
	},
	{
		name:     "integer precision 2 format",
		line:     testLine(),
		bytes:    0,
		format:   "%-04.2d",
		expected: "00B ",
	},
	{
		name:     "integer precision 2 format",
		line:     testLine(),
		bytes:    0,
		format:   "%+.2d",
		expected: "+00B",
	},
	{
		name:     "integer precision 2 format",
		line:     testLine(),
		bytes:    0,
		format:   "%#.2d",
		expected: "00B",
	},
	{
		name:     "integer precision 2 format",
		line:     testLine(),
		bytes:    1,
		format:   "%.2d",
		expected: "01B",
	},
	{
		name:      "integer precision 2 format",
		line:      testLine(),
		bytes:     1,
		format:    "% .2d",
		expected:  "01 B",
		benchmark: true,
	},
	{
		name:     "integer precision 2 format",
		line:     testLine(),
		bytes:    1,
		format:   "% 2.2d",
		expected: "01  B",
	},
	{
		name:     "integer precision 2 format",
		line:     testLine(),
		bytes:    1,
		format:   "%2.2d",
		expected: "01B",
	},
	{
		name:     "integer precision 2 format",
		line:     testLine(),
		bytes:    1,
		format:   "%4.2d",
		expected: " 01B",
	},
	{
		name:     "integer precision 2 format",
		line:     testLine(),
		bytes:    1,
		format:   "%03.2d",
		expected: "01B",
	},
	{
		name:     "integer precision 2 format",
		line:     testLine(),
		bytes:    1,
		format:   "%04.2d",
		expected: "001B",
	},
	{
		name:     "integer precision 2 format",
		line:     testLine(),
		bytes:    1,
		format:   "%-3.2d",
		expected: "01B",
	},
	{
		name:     "integer precision 2 format",
		line:     testLine(),
		bytes:    1,
		format:   "%-4.2d",
		expected: "01B ",
	},
	{
		name:     "integer precision 2 format",
		line:     testLine(),
		bytes:    1,
		format:   "%-03.2d",
		expected: "01B",
	},
	{
		name:     "integer precision 2 format",
		line:     testLine(),
		bytes:    1,
		format:   "%-04.2d",
		expected: "01B ",
	},
	{
		name:     "integer precision 2 format",
		line:     testLine(),
		bytes:    1,
		format:   "%+.2d",
		expected: "+01B",
	},
	{
		name:     "integer precision 2 format",
		line:     testLine(),
		bytes:    1,
		format:   "%#.2d",
		expected: "01B",
	},
	{
		name:     "integer precision 2 format",
		line:     testLine(),
		bytes:    1024,
		format:   "%.2d",
		expected: "01K",
	},
	{
		name:      "integer precision 2 format",
		line:      testLine(),
		bytes:     1024,
		format:    "% .2d",
		expected:  "01 K",
		benchmark: true,
	},
	{
		name:     "integer precision 2 format",
		line:     testLine(),
		bytes:    1024,
		format:   "% 2.2d",
		expected: "01  K",
	},
	{
		name:     "integer precision 2 format",
		line:     testLine(),
		bytes:    1024,
		format:   "%3.2d",
		expected: "01K",
	},
	{
		name:     "integer precision 2 format",
		line:     testLine(),
		bytes:    1024,
		format:   "%4.2d",
		expected: " 01K",
	},
	{
		name:     "integer precision 2 format",
		line:     testLine(),
		bytes:    1024,
		format:   "%03.2d",
		expected: "01K",
	},
	{
		name:     "integer precision 2 format",
		line:     testLine(),
		bytes:    1024,
		format:   "%04.2d",
		expected: "001K",
	},
	{
		name:     "integer precision 2 format",
		line:     testLine(),
		bytes:    1024,
		format:   "%-3.2d",
		expected: "01K",
	},
	{
		name:     "integer precision 2 format",
		line:     testLine(),
		bytes:    1024,
		format:   "%-4.2d",
		expected: "01K ",
	},
	{
		name:     "integer precision 2 format",
		line:     testLine(),
		bytes:    1024,
		format:   "%-03.2d",
		expected: "01K",
	},
	{
		name:     "integer precision 2 format",
		line:     testLine(),
		bytes:    1024,
		format:   "%-04.2d",
		expected: "01K ",
	},
	{
		name:     "integer precision 2 format",
		line:     testLine(),
		bytes:    1024,
		format:   "%+.2d",
		expected: "+01K",
	},
	{
		name:     "integer precision 2 format",
		line:     testLine(),
		bytes:    1024,
		format:   "%#.2d",
		expected: "01K",
	},
	{
		name:     "integer precision 2 format",
		line:     testLine(),
		bytes:    1128,
		format:   "%.2d",
		expected: "01K",
	},
	{
		name:      "integer precision 2 format",
		line:      testLine(),
		bytes:     1128,
		format:    "% .2d",
		expected:  "01 K",
		benchmark: true,
	},
	{
		name:     "integer precision 2 format",
		line:     testLine(),
		bytes:    1128,
		format:   "% 2.2d",
		expected: "01  K",
	},
	{
		name:     "integer precision 2 format",
		line:     testLine(),
		bytes:    1128,
		format:   "%3.2d",
		expected: "01K",
	},
	{
		name:     "integer precision 2 format",
		line:     testLine(),
		bytes:    1128,
		format:   "%4.2d",
		expected: " 01K",
	},
	{
		name:     "integer precision 2 format",
		line:     testLine(),
		bytes:    1128,
		format:   "%02.2d",
		expected: "01K",
	},
	{
		name:     "integer precision 2 format",
		line:     testLine(),
		bytes:    1128,
		format:   "%04.2d",
		expected: "001K",
	},
	{
		name:     "integer precision 2 format",
		line:     testLine(),
		bytes:    1128,
		format:   "%-2.2d",
		expected: "01K",
	},
	{
		name:     "integer precision 2 format",
		line:     testLine(),
		bytes:    1128,
		format:   "%-4.2d",
		expected: "01K ",
	},
	{
		name:     "integer precision 2 format",
		line:     testLine(),
		bytes:    1128,
		format:   "%-02.2d",
		expected: "01K",
	},
	{
		name:     "integer precision 2 format",
		line:     testLine(),
		bytes:    1128,
		format:   "%-04.2d",
		expected: "01K ",
	},
	{
		name:     "integer precision 2 format",
		line:     testLine(),
		bytes:    1128,
		format:   "%+.2d",
		expected: "+01K",
	},
	{
		name:     "integer precision 2 format",
		line:     testLine(),
		bytes:    1128,
		format:   "%#.2d",
		expected: "01K",
	},
	{
		name:     "name",
		line:     testLine(),
		bytes:    1124,
		format:   "%d",
		names:    []string{"B", "Kilobyte"},
		expected: "1Kilobyte",
	},
	{
		name:      "name",
		line:      testLine(),
		bytes:     1124,
		format:    "% d",
		names:     []string{"B", "Kilobyte"},
		expected:  "1 Kilobyte",
		benchmark: true,
	},
	{
		name:     "name",
		line:     testLine(),
		bytes:    1124,
		format:   "% 2d",
		names:    []string{"B", "Kilobyte"},
		expected: "1  Kilobyte",
	},
}

func TestBytesFormat(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	for _, tc := range BytesFormatTestCases {
		tc := tc
		t.Run(tc.name+" "+tc.format+" "+strconv.FormatUint(tc.bytes, 10)+" "+strconv.Itoa(tc.line), func(t *testing.T) {
			t.Parallel()
			linkToExample := fmt.Sprintf("%s:%d", testFile, tc.line)
			s := fmt.Sprintf(tc.format, New(tc.bytes, tc.names...))
			if s != tc.expected {
				t.Errorf("unexpected kilobytes %#v, expected %#v %s",
					s, tc.expected, linkToExample)
			}
		})
	}
}

func BenchmarkBytesFormat(b *testing.B) {
	b.ReportAllocs()
	for _, tc := range BytesFormatTestCases {
		if !tc.benchmark {
			continue
		}
		b.Run(tc.name+" "+tc.format+" "+strconv.FormatUint(tc.bytes, 10)+" "+strconv.Itoa(tc.line), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = fmt.Sprintf(tc.format, New(tc.bytes, tc.names...))
			}
		})
	}
}
