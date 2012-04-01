// Copyright 2012 Harry de Boer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package matrix

// MulStrassen returns A * B.
//
// Original paper: Gaussian Elimination is not Optimal.
//                 Volker Strassen, 1969.
//
// This implementation is not optimized, it serves as a reference for testing.
func MulStrassen(A, B *Matrix) *Matrix {
	return Zeros(A.height, B.width).MulStrassen(A, B)
}

// MulStrassen calculates C = A * B and returs C.
func (C *Matrix) MulStrassen(A, B *Matrix) *Matrix {

	if A.width < 80 || A.height != A.width || A.height % 2 != 0 {
		return C.MulBLAS(A, B)
	}

	m := A.height / 2
	A11 := A.SubMatrix(0, 0, m, m)
	A12 := A.SubMatrix(0, m, m, m)
	A21 := A.SubMatrix(m, 0, m, m)
	A22 := A.SubMatrix(m, m, m, m)
	B11 := B.SubMatrix(0, 0, m, m)
	B12 := B.SubMatrix(0, m, m, m)
	B21 := B.SubMatrix(m, 0, m, m)
	B22 := B.SubMatrix(m, m, m, m)
	C11 := C.SubMatrix(0, 0, m, m)
	C12 := C.SubMatrix(0, m, m, m)
	C21 := C.SubMatrix(m, 0, m, m)
	C22 := C.SubMatrix(m, m, m, m)

	M1 := MulStrassen(Plus(A11, A22), Plus(B11, B22))
	M2 := MulStrassen(Plus(A21, A22), B11)
	M3 := MulStrassen(A11, Minus(B12, B22))
	M4 := MulStrassen(A22, Minus(B21, B11))
	M5 := MulStrassen(Plus(A11, A12), B22)
	M6 := MulStrassen(Minus(A21, A11), Plus(B11, B12))
	M7 := MulStrassen(Minus(A12, A22), Plus(B21, B22))

	C11.Add(M7).Add(M1).Add(M4).Sub(M5)
	C12.Add(M5).Add(M3)
	C21.Add(M4).Add(M2)
	C22.Add(M6).Add(M1).Sub(M2).Add(M3)
	return C
}
