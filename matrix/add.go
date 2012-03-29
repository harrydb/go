// Copyright 2012 Harry de Boer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package matrix

// Plus returns A + B.
func Plus(A, B *Matrix) *Matrix {
	 C := Zeros(A.height, A.width)
	 C.Plus(A, B)
	return C
}

// Add calculates A = A + B.
func (A *Matrix) Add(B *Matrix) {

	// Normal matrices.
	if A.stride == A.width && B.stride == B.width {
		for i, bi := range B.data {
			A.data[i] += bi
		}
		return
	}

	// Submatrices.
	for i := 0; i < A.height; i++ {
		Ai := A.Row(i)
		for j, bij := range B.Row(i) {
			Ai[j] += bij
		}
	}
}

// Plus calculates C = A + B.
func (C *Matrix) Plus(A, B *Matrix) {

	// Normal matrices.
	if A.stride == A.width && B.stride == B.width {
		for i, ai := range A.data {
			C.data[i] = ai + B.data[i]
		}
		return
	}

	// SubMatrices.
	for i, k := 0, 0; i < A.height; i++ {
		Ai := A.Row(i)
		for j, bij := range B.Row(i) {
			C.data[k] = Ai[j] + bij
			k++
		}
	}
}
