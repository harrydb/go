// Copyright 2012 Harry de Boer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package matrix

// MulSimple returns A * B.
func MulSimple(A, B *Matrix) *Matrix {
	C := Zeros(A.height, B.width)
	C.MulSimple(A, B)
	return C
}

// MulSimple calculates C = A * B.
func (C *Matrix) MulSimple(A, B *Matrix) {
	for i := 0; i < C.height; i++ {
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
}

// MulAddSimple calculates C = C + A * B.
func (C *Matrix) MulAddSimple(A, B *Matrix) {
	for i := 0; i < C.height; i++ {
		Ci := C.Row(i)

		for j, aij := range A.Row(i) {
			for k, bjk := range B.Row(j) {
				Ci[k] += aij * bjk
			}
		}
	}
}

// MulSubSimple calculates C = C - A * B.
func (C *Matrix) MulSubSimple(A, B *Matrix) {
	for i := 0; i < C.height; i++ {
		Ci := C.Row(i)

		for j, aij := range A.Row(i) {
			for k, bjk := range B.Row(j) {
				Ci[k] -= aij * bjk
			}
		}
	}
}
