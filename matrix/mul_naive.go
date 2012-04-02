// Copyright 2012 Harry de Boer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package matrix

// MulNaive returns A * B.
func MulNaive(A, B *Matrix) *Matrix {
	return Zeros(A.height, B.width).MulNaive(A, B)
}

// MulNaive calculates C = A * B and returns C.
func (C *Matrix) MulNaive(A, B *Matrix) *Matrix {
	for i := 0; i < A.height; i++ {
		Ci := C.Row(i)
		for k := range Ci {
			Ci[k] = 0
		}
		for j, aij := range A.Row(i) {
			for k, bjk := range B.Row(j) {
				Ci[k] += aij * bjk
			}
		}
	}
	return C
}

// MulAddNaive calculates C = C + A * B and returns C.
func (C *Matrix) MulAddNaive(A, B *Matrix) *Matrix{
	for i := 0; i < A.height; i++ {
		Ci := C.Row(i)

		for j, aij := range A.Row(i) {
			for k, bjk := range B.Row(j) {
				Ci[k] += aij * bjk
			}
		}
	}
	return C
}

// MulSubNaive calculates C = C - A * B and returns C.
func (C *Matrix) MulSubNaive(A, B *Matrix) *Matrix {
	for i := 0; i < A.height; i++ {
		Ci := C.Row(i)
		for j, aij := range A.Row(i) {
			for k, bjk := range B.Row(j) {
				Ci[k] -= aij * bjk
			}
		}
	}
	return C
}
