// Copyright 2012 Harry de Boer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package matrix

// MulStrassen returns C = A * B.
// See Strassen, 1969.
func MulStrassen(A, B *Matrix) *Matrix {
	return Zeros(A.height, A.width).MulAddStrassen(A, B)
}

// MulAddStrassen returns C = C + A * B.
func (C *Matrix) MulAddStrassen(A, B *Matrix) *Matrix {

	if A.width < 80 || A.height != A.width || A.height % 2 != 0 {
		return C.MulAddBLAS(A, B)
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

	M1 := Zeros(m, m)
	M2 := Zeros(m, m)
	M3 := Zeros(m, m)
	M1.MulAddStrassen(Plus(A11, A22), Plus(B11, B22))
	M2.MulAddStrassen(Plus(A21, A22), B11)
	M3.MulAddStrassen(A11, Minus(B12, B22))
	C21.MulAddStrassen(A22, Minus(B21, B11))
	C12.MulAddStrassen(Plus(A11, A12), B22)
	C22.MulAddStrassen(Minus(A21, A11), Plus(B11, B12))
	C11.MulAddStrassen(Minus(A12, A22), Plus(B21, B22))

	C11.Add(M1).Add(C21).Sub(C12);
	C12.Add(M3);
	C21.Add(M2);
	C22.Add(M1).Sub(M2).Add(M3);
	return C
}
