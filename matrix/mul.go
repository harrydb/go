// Copyright 2012 Harry de Boer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package matrix

import "runtime"

// Mul returns A * B.
func Mul(A, B *Matrix) *Matrix {
	n := (A.height / 2) + (A.width / 2)

	if runtime.GOMAXPROCS(0) > 1 {
		switch {
			case n < 32: return MulBLAS(A, B)
			default: return MulStrassenPar(A, B)
		}
	}

	if n < 80 {
		return MulBLAS(A, B)
	}
	return MulDouglas(A, B)
}

// Mul calculates C = A * B and returns C.
func (C *Matrix) Mul(A, B *Matrix) *Matrix {
	return C.MulBLAS(A, B)
}

// MulAdd calculates C = C + A * B and returns C.
func (C *Matrix) MulAdd(A, B *Matrix) *Matrix{
	return C.MulAddBLAS(A, B)
}

// MulSub calculates C = C - A * B and returns C.
func (C *Matrix) MulSub(A, B *Matrix) *Matrix {
	return C.MulSubBLAS(A, B)
}
