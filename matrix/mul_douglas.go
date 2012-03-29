// Copyright 2012 Harry de Boer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package matrix

// MulDouglas returns A * B.
// See Douglas et al, 1994.
func MulDouglas(A, B *Matrix) *Matrix {
	C := Zeros(A.height, B.width)
	C.MulDouglas(A, B)
	return C
}

func (C *Matrix) MulDouglas(A, B *Matrix) {

	if A.width < 80 || A.height != A.width || A.height % 2 != 0 {
		C.MulBLAS(A, B)
		return
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

	X := Minus(A11, A21)
	Y := Minus(B22, B12)
	C21.MulDouglas(X, Y)
	X.Plus(A21, A22)
	Y.Minus(B12, B11)
	C22.MulDouglas(X, Y)
	X.SubBLAS(A11)
	Y.Minus(B22, Y)
	C12.MulDouglas(X, Y)
	X.Minus(A12, X)
	C11.MulDouglas(X, B22)
	X.MulDouglas(A11, B11)
	C12.AddBLAS(X)
	C21.AddBLAS(C12)
	C12.AddBLAS(C22)
	C22.AddBLAS(C21) // final c22
	C12.AddBLAS(C11) // final c12
	Y.SubBLAS(B21)
	C11.MulDouglas(A22, Y)
	C21.SubBLAS(C11) // final c21
	C11.MulDouglas(A12, B21)
	C11.AddBLAS(X) // final c11
}
