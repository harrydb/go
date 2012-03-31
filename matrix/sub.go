// Copyright 2012 Harry de Boer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package matrix

// Minus returns A - B.
func Minus(A, B *Matrix) *Matrix {
	 C := Zeros(A.height, A.width)
	 C.Minus(A, B)
	return C
}

// Subtract calculates A = A - B and returns A.
func (A *Matrix) Sub(B *Matrix) *Matrix {

	// Normal matrices.
	if A.stride == A.width && B.stride == B.width {
		for i, bi := range B.data {
			A.data[i] -= bi
		}
		return A
	}

	// Submatrices.
	for i := 0; i < A.height; i++ {
		Ai := A.Row(i)
		for j, bij := range B.Row(i) {
			Ai[j] -= bij
		}
	}
	return A
}

// Minus calculates C = A - B and returns C.
func (C *Matrix) Minus(A, B *Matrix) *Matrix {

	// Normal matrices.
	if A.stride == A.width && B.stride == B.width {
		for i, ai := range A.data {
			C.data[i] = ai - B.data[i]
		}
		return C
	}

	// SubMatrices.
	for i, k := 0, 0; i < A.height; i++ {
		Ai := A.Row(i)
		for j, bij := range B.Row(i) {
			C.data[k] = Ai[j] - bij
			k++
		}
	}
	return C
}
