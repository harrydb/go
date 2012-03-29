// Copyright 2012 Harry de Boer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package matrix

import "github.com/ziutek/blas"

// MulBLAS returns A * B.
func MulBLAS(A, B *Matrix) *Matrix {
	C := Zeros(A.height, B.width)
	C.MulAddBLAS(A, B)
	return C
}

// MulBLAS calculates C = A * B.
func (C *Matrix) MulBLAS(A, B *Matrix) {
	for i := 0; i < C.height; i++ {
		Ci := C.Row(i)
		for k := range Ci {
			Ci[k] = 0
		}
		for j, aij := range A.Row(i) {
			// Ci = aij * Bj
			blas.Daxpy(C.width, aij, B.Row(j), 1, Ci, 1)
		}
	}
}

// MulAddBLAS calculates C = C + A * B.
func (C *Matrix) MulAddBLAS(A, B *Matrix) {
	for i := 0; i < C.height; i++ {
		Ci := C.Row(i)
		for j, aij := range A.Row(i) {
			// Ci = Ci + aij * Bj
			blas.Daxpy(C.width, aij, B.Row(j), 1, Ci, 1)
		}
	}
}

// MulSubBLAS calculates C = C - A * B.
func (C *Matrix) MulSubBLAS(A, B *Matrix) {
	for i := 0; i < C.height; i++ {
		Ci := C.Row(i)
		for j, aij := range A.Row(i) {
			// Ci = Ci - aij * Bj
			blas.Daxpy(C.width, -aij, B.Row(j), 1, Ci, 1)
		}
	}
}

