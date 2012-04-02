// Copyright 2012 Harry de Boer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package matrix

func MulHuss(A, B *Matrix) *Matrix {
	return Zeros(A.height, B.width).MulAddHuss(A, B)
}

// MulAddHuss returns C = C + A * B.
func (C *Matrix) MulAddHuss(A, B *Matrix) *Matrix {

	if A.width < 80 || A.height != A.width || A.height % 2 != 0 {
		return C.MulAddBLAS(A, B)
	}

	n := A.height / 2
	A11 := A.SubMatrix(0, 0, n, n)
	A12 := A.SubMatrix(0, n, n, n)
	A21 := A.SubMatrix(n, 0, n, n)
	A22 := A.SubMatrix(n, n, n, n)
	B11 := B.SubMatrix(0, 0, n, n)
	B12 := B.SubMatrix(0, n, n, n)
	B21 := B.SubMatrix(n, 0, n, n)
	B22 := B.SubMatrix(n, n, n, n)
	C11 := C.SubMatrix(0, 0, n, n)
	C12 := C.SubMatrix(0, n, n, n)
	C21 := C.SubMatrix(n, 0, n, n)
	C22 := C.SubMatrix(n, n, n, n)

	// Allocate scratch space
	X := Zeros(n, n)
	Y := Zeros(n, n)
	Z := Zeros(n, n)

	// Perform calculations.
	X.PlusBLAS(A21, A22)
	Y.Minus(B12, B11)
	Z.MulAddHuss(X, Y)
	C22.AddBLAS(Z)
	C12.AddBLAS(Z)
	X.SubBLAS(A11)
	Y.Minus(B22, Y)
	Z.Clear()
	Z.MulAddHuss(A11, B11)
	C11.AddBLAS(Z)
	Z.MulAddHuss(X, Y)
	C11.MulAddHuss(A12, B21) // final C11
	X.Minus(A12, X)
	Y.SubBLAS(B21)
	C12.MulAddHuss(X, B22)
	C12.AddBLAS(Z) // final C12
	C21.ScaleBLAS(-1)
	C21.MulAddHuss(A22, Y)
	X.Minus(A11, A21)
	Y.Minus(B22, B12)
	Z.MulAddHuss(X, Y)
	C22.AddBLAS(Z) // final C22
	C21.Minus(Z, C21) // final C21
	return C
}
