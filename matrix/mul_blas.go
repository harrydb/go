// Copyright 2012 Harry de Boer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package matrix

import "github.com/ziutek/blas"

// MulBLAS returns A * B.
//
// Performance of this implementation is improved by using level 1 BLAS functions.
func MulBLAS(A, B *Matrix) *Matrix {
	return Zeros(A.height, B.width).MulAddBLAS(A, B)
}

// MulBLAS calculates C = A * B.
func (C *Matrix) MulBLAS(A, B *Matrix) *Matrix {
	for i := 0; i < A.height; i++ {
		Ci := C.Row(i)
		for k := range Ci {
			Ci[k] = 0
		}
		for j, aij := range A.Row(i) {
			// Ci += aij * Bj
			blas.Daxpy(C.width, aij, B.Row(j), 1, Ci, 1)
		}
	}
	return C
}

// MulAddBLAS calculates C = C + A * B.
func (C *Matrix) MulAddBLAS(A, B *Matrix) *Matrix {
	for i := 0; i < A.height; i++ {
		Ci := C.Row(i)
		for j, aij := range A.Row(i) {
			// Ci = Ci + aij * Bj
			blas.Daxpy(C.width, aij, B.Row(j), 1, Ci, 1)
		}
	}
	return C
}

// MulSubBLAS calculates C = C - A * B.
func (C *Matrix) MulSubBLAS(A, B *Matrix) *Matrix {
	for i := 0; i < A.height; i++ {
		Ci := C.Row(i)
		for j, aij := range A.Row(i) {
			// Ci = Ci - aij * Bj
			blas.Daxpy(C.width, -aij, B.Row(j), 1, Ci, 1)
		}
	}
	return C
}

