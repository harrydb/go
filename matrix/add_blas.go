// Copyright 2012 Harry de Boer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package matrix

import "github.com/ziutek/blas"

// Plus returns A + B.
func PlusBLAS(A, B *Matrix) *Matrix {
	return Zeros(A.height, A.width).PlusBLAS(A, B)
}



// Add calculates A = A + B and returns A.
func (A *Matrix) AddBLAS(B *Matrix) *Matrix {

	// Normal matrices.
	if A.stride == A.width && B.stride == B.width {
		blas.Daxpy(len(A.data), 1.0, B.data, 1, A.data, 1)
		return A
	}

	// Submatrices.
	for i := 0; i < A.height; i++ {
		blas.Daxpy(A.width, 1.0, B.Row(i), 1, A.Row(i), 1)
	}

	return A
}


// Copy the contents of B to A.
func (A *Matrix) CopyBLAS(B *Matrix) {

	// Normal matrices.
	if B.stride == B.width {
		blas.Dcopy(len(A.data), B.data, 1, A.data, 1)
		return
	}

	// Submatrices.
	for i := 0; i < A.height; i++ {
		blas.Dcopy(A.width, B.Row(i), 1, A.Row(i), 1)
	}
}

// Plus calculates C = A + B and returns C.
func (C *Matrix) PlusBLAS(A, B *Matrix) *Matrix {
	if C != B {
		C.CopyBLAS(B)
	}
	return C.AddBLAS(A)
}
