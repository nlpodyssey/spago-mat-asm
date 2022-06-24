// Copyright 2022 The NLP Odyssey Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build amd64 && gc && !purego

package matfuncs

// AddConst32 adds a constant value c to each element of x, storing the result in y (32 bits).
func AddConst32(c float32, x, y []float32) {
	if hasAVX2 {
		AddConstAVX32(c, x, y)
		return
	}
	AddConstSSE32(c, x, y)
}

// AddConst64 adds a constant value c to each element of x, storing the result in y (64 bits).
func AddConst64(c float64, x, y []float64) {
	if hasAVX2 {
		AddConstAVX64(c, x, y)
		return
	}
	AddConstSSE64(c, x, y)
}
