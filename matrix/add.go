// Copyright 2012 Harry de Boer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package matrix

// Plus returns A + B.
func Plus(A, B *Matrix) *Matrix {
	 return Zeros(A.height, A.width).Plus(A, B)
}

// Add calculates A = A + B and returns A.
func (A *Matrix) Add(B *Matrix) *Matrix {

	// Normal matrices.
	if A.stride == A.width && B.stride == B.width {
		for i, bi := range B.data {
			A.data[i] += bi
		}
		return A
	}

	// Submatrices.
	for i := 0; i < A.height; i++ {
		Ai := A.Row(i)
		for j, bij := range B.Row(i) {
			Ai[j] += bij
		}
	}

	return A
}

// Plus calculates C = A + B and returns C.
func (C *Matrix) Plus(A, B *Matrix) *Matrix  {

	// Normal matrices.
	if A.stride == A.width && B.stride == B.width {
		for i, ai := range A.data {
			C.data[i] = ai + B.data[i]
		}
		return C
	}

	// SubMatrices.
	for i, k := 0, 0; i < A.height; i++ {
		Ai := A.Row(i)
		for j, bij := range B.Row(i) {
			C.data[k] = Ai[j] + bij
			k++
		}
	}

	return C
}
