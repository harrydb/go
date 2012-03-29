// Copyright 2012 Harry de Boer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package matrix

import "github.com/ziutek/blas"

func MinusBLAS(A, B *Matrix) *Matrix {
	 C := Zeros(A.height, A.width)
	 C.MinusBLAS(A, B)
	return C
}

// Subtract calculates A = A - B.
func (A *Matrix) SubBLAS(B *Matrix) {

	// Normal matrices.
	if A.stride == A.width && B.stride == B.width {
		blas.Daxpy(len(A.data), -1.0, B.data, 1, A.data, 1)
		return
	}

	// Submatrices.
	for i := 0; i < A.height; i++ {
		blas.Daxpy(A.width, -1.0, B.Row(i), 1, A.Row(i), 1)
	}
}


// Minus calculates C = A - B.
func (C *Matrix) MinusBLAS(A, B *Matrix) {

	if C == B {
		C.ScaleBLAS(-1)
		C.AddBLAS(A)
		return
	}

	C.CopyBLAS(A)
	C.SubBLAS(B)
}
