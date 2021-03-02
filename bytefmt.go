// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package bytefmt provides Bytes type whats implements fmt.Formatter interface.
package bytefmt

import (
	"fmt"
	"math"
	"strconv"
)

const (
	Byte uint64 = 1 << (10 * iota)
	Kilobyte
	Megabyte
	Gigabyte
	Terabyte
	Petabyte
	Exabyte
)

type Bytes struct {
	Value uint64
	names []string
}

func New(v uint64, n ...string) Bytes {
	b := Bytes{Value: v}
	b.Names(n...)
	return b
}

var names = []string{"B", "K", "M", "G", "T", "P", "E"}

func (b *Bytes) Names(n ...string) []string {
	if len(n) == 0 {
		if len(b.names) == 0 {
			b.names = names
		}
		return b.names
	}
	b.names = n
	if len(b.names) < len(names) {
		b.names = append(b.names, names[len(b.names):]...)
	}
	return b.names
}

func (b Bytes) String() string {
	var (
		f float64
		i int
	)
	if b.Value >= Exabyte {
		f = b.float64()
		i = 6
	} else if b.Value >= Petabyte {
		f = b.float64()
		i = 5
	} else if b.Value >= Terabyte {
		f = b.float64()
		i = 4
	} else if b.Value >= Gigabyte {
		f = b.float64()
		i = 3
	} else if b.Value >= Megabyte {
		f = b.float64()
		i = 2
	} else if b.Value >= Kilobyte {
		f = b.float64()
		i = 1
	} else {
		f = b.float64()
		i = 0
	}
	return fmt.Sprintf("%v%s", f, names[i])
}

/*
	Other flags:
		+	always print a sign for numeric values (%+q);
		-	pad with spaces on the right rather than the left (left-justify the field)
		' '	(space) leave a space between value and units (% d);
		0	pad with leading zeros rather than spaces;
*/

func (b Bytes) Format(f fmt.State, c rune) {
	var (
		v  interface{}      // value
		vf = "%"            // value format
		uf = "%"            // unit format
		pf = ""             // padding format
		u  = b.name()       // unit of measure name
		uw = len([]rune(u)) // unit name width
	)
	if f.Flag('+') {
		vf += "+"
	}
	if f.Flag('#') {
		vf += "#"
	}
	w, ok := f.Width()
	if !ok {
		w = -1
	}
	if f.Flag(' ') {
		if w == -1 {
			uf += strconv.Itoa(uw + 1)
		} else {
			uf += strconv.Itoa(uw + w)
		}
	} else if w != -1 {
		pf += "%"
		if f.Flag('-') {
			pf += "-"
		} else if f.Flag('0') {
			pf += "0"
		}
		pf += strconv.Itoa(w) + "s"
	}
	if p, ok := f.Precision(); ok {
		vf += "." + strconv.Itoa(p)
	}
	vf += string(c)
	uf += "s"
	var s string
	if c == 's' || c == 'q' {
		v = fmt.Sprintf("%v", b.float64())
		s = fmt.Sprint(v) + fmt.Sprintf(uf, u)
		if pf != "" {
			s = fmt.Sprintf(pf, s)
		}
		s = fmt.Sprintf(vf, s)
	} else {
		if c == 'd' {
			v = int64(math.Round(b.float64()))
		} else {
			v = b.float64()
		}
		s = fmt.Sprintf(vf+uf, v, u)
		if pf != "" {
			s = fmt.Sprintf(pf, s)
		}
	}
	f.Write([]byte(s))
}

// float64 returns returns a number in units of measure
// when converted to which the smallest integer is obtained
func (b Bytes) float64() float64 {
	if b.Value >= Exabyte {
		return b.exabytes()
	} else if b.Value >= Petabyte {
		return b.petabytes()
	} else if b.Value >= Terabyte {
		return b.terabytes()
	} else if b.Value >= Gigabyte {
		return b.gigabytes()
	} else if b.Value >= Megabyte {
		return b.megabytes()
	} else if b.Value >= Kilobyte {
		return b.kilobytes()
	} else {
		return float64(b.Value)
	}
}

func (b Bytes) kilobytes() float64 { return float64(b.Value) / float64(Kilobyte) }
func (b Bytes) megabytes() float64 { return float64(b.Value) / float64(Megabyte) }
func (b Bytes) gigabytes() float64 { return float64(b.Value) / float64(Gigabyte) }
func (b Bytes) terabytes() float64 { return float64(b.Value) / float64(Terabyte) }
func (b Bytes) petabytes() float64 { return float64(b.Value) / float64(Petabyte) }
func (b Bytes) exabytes() float64  { return float64(b.Value) / float64(Exabyte) }

// name returns the name of the unit of measure
// when converted to which the smallest integer is obtained
func (b Bytes) name() string {
	if b.Value >= Exabyte {
		return b.names[6]
	} else if b.Value >= Petabyte {
		return b.names[5]
	} else if b.Value >= Terabyte {
		return b.names[4]
	} else if b.Value >= Gigabyte {
		return b.names[3]
	} else if b.Value >= Megabyte {
		return b.names[2]
	} else if b.Value >= Kilobyte {
		return b.names[1]
	} else {
		return b.names[0]
	}
}
