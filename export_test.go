// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bytefmt

// Exported for testing only.

func (b Bytes) Kilobytes() float64 { return b.kilobytes() }
func (b Bytes) Megabytes() float64 { return b.megabytes() }
func (b Bytes) Gigabytes() float64 { return b.gigabytes() }
func (b Bytes) Terabytes() float64 { return b.terabytes() }
func (b Bytes) Petabytes() float64 { return b.petabytes() }
func (b Bytes) Exabytes() float64  { return b.exabytes() }
