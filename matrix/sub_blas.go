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

// Subtract calculates A = A - B and returns A.
func (A *Matrix) SubBLAS(B *Matrix) *Matrix {

	// Normal matrices.
	if A.stride == A.width && B.stride == B.width {
		blas.Daxpy(len(A.data), -1.0, B.data, 1, A.data, 1)
		return A
	}

	// Submatrices.
	for i := 0; i < A.height; i++ {
		blas.Daxpy(A.width, -1.0, B.Row(i), 1, A.Row(i), 1)
	}
	return A
}


// Minus calculates C = A - B and returns C.
func (C *Matrix) MinusBLAS(A, B *Matrix) *Matrix {

	if C == B {
		C.ScaleBLAS(-1)
		return C.AddBLAS(A)
	}

	C.Copy(A)
	return C.SubBLAS(B)
}
