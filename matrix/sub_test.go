// Copyright 2012 Harry de Boer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package matrix

import (
	"math/rand"
	"testing"
	"time"
)

func TestMinus(t *testing.T) {
	A := New(2, 3, []float64{1, 2, 4, 2, 1, 2})
	B := New(2, 3, []float64{8, 4, 2, 2, 1, 0})
	C := Minus(A, B)

	correct := New(2, 3, []float64{-7, -2, 2, 0, 0, 2})
	if !equal(C, correct, 0, t) {
		t.FailNow()
	}
}

func TestMinusBLAS(t *testing.T) {
	A := New(2, 3, []float64{1, 2, 4, 2, 1, 2})
	B := New(2, 3, []float64{8, 4, 2, 2, 1, 0})
	C := MinusBLAS(A, B)

	correct := New(2, 3, []float64{-7, -2, 2, 0, 0, 2})
	if !equal(C, correct, 0, t) {
		t.FailNow()
	}
}

func TestSub(t *testing.T) {
	A := New(2, 3, []float64{1, 2, 4, 2, 1, 2})
	B := New(2, 3, []float64{8, 4, 2, 2, 1, 0})
	A.Sub(B)

	correct := New(2, 3, []float64{-7, -2, 2, 0, 0, 2})
	if !equal(A, correct, 0, t) {
		t.FailNow()
	}
}

func TestSubBLAS(t *testing.T) {
	A := New(2, 3, []float64{1, 2, 4, 2, 1, 2})
	B := New(2, 3, []float64{8, 4, 2, 2, 1, 0})
	A.SubBLAS(B)

	correct := New(2, 3, []float64{-7, -2, 2, 0, 0, 2})
	if !equal(A, correct, 0, t) {
		t.FailNow()
	}
}

func TestSubSubmatrix(t *testing.T) {
	rand.Seed(time.Now().Unix())
	n := 4
	A := randomMatrix(n, n)
	B := randomMatrix(n, n)
	C := A.SubMatrix(2, 2, 2, 2)
	D := B.SubMatrix(2, 2, 2, 2)
	C.Sub(D)
}

func BenchmarkSubBLAS_____1024(bench *testing.B) {
	bench.StopTimer()
	rand.Seed(time.Now().Unix())
	n := 1024
	A := randomMatrix(n, n)
	B := randomMatrix(n, n)
	bench.StartTimer()
	for i := 0; i < bench.N; i++ {
		A.SubBLAS(B)
    }
}

func BenchmarkSub_________1024(bench *testing.B) {
	bench.StopTimer()
	rand.Seed(time.Now().Unix())
	n := 1024
	A := randomMatrix(n, n)
	B := randomMatrix(n, n)
	bench.StartTimer()
	for i := 0; i < bench.N; i++ {
		A.Sub(B)
    }
}

func BenchmarkMinusSBLAS___1024(bench *testing.B) {
	bench.StopTimer()
	rand.Seed(time.Now().Unix())
	n := 1024
	A := randomMatrix(n, n)
	B := randomMatrix(n, n)
	C := Zeros(n, n)
	bench.StartTimer()
	for i := 0; i < bench.N; i++ {
		C.MinusBLAS(A, B)
    }
}

func BenchmarkMinusS_______1024(bench *testing.B) {
	bench.StopTimer()
	rand.Seed(time.Now().Unix())
	n := 1024
	A := randomMatrix(n, n)
	B := randomMatrix(n, n)
	C := Zeros(n, n)
	bench.StartTimer()
	for i := 0; i < bench.N; i++ {
		C.Minus(A, B)
    }
}

func BenchmarkMinusBLAS____1024(bench *testing.B) {
	bench.StopTimer()
	rand.Seed(time.Now().Unix())
	n := 1024
	A := randomMatrix(n, n)
	B := randomMatrix(n, n)
	bench.StartTimer()
	for i := 0; i < bench.N; i++ {
		MinusBLAS(A, B)
    }
}

func BenchmarkMinus________1024(bench *testing.B) {
	bench.StopTimer()
	rand.Seed(time.Now().Unix())
	n := 1024
	A := randomMatrix(n, n)
	B := randomMatrix(n, n)
	bench.StartTimer()
	for i := 0; i < bench.N; i++ {
		Minus(A, B)
    }
}

func BenchmarkSubBLAS______256(bench *testing.B) {
	bench.StopTimer()
	rand.Seed(time.Now().Unix())
	n := 256
	A := randomMatrix(n, n)
	B := randomMatrix(n, n)
	bench.StartTimer()
	for i := 0; i < bench.N; i++ {
		A.SubBLAS(B)
    }
}

func BenchmarkSub__________256(bench *testing.B) {
	bench.StopTimer()
	rand.Seed(time.Now().Unix())
	n := 256
	A := randomMatrix(n, n)
	B := randomMatrix(n, n)
	bench.StartTimer()
	for i := 0; i < bench.N; i++ {
		A.Sub(B)
    }
}

func BenchmarkMinusBLAS_____256(bench *testing.B) {
	bench.StopTimer()
	rand.Seed(time.Now().Unix())
	n := 256
	A := randomMatrix(n, n)
	B := randomMatrix(n, n)
	bench.StartTimer()
	for i := 0; i < bench.N; i++ {
		MinusBLAS(A, B)
    }
}

func BenchmarkMinus_________256(bench *testing.B) {
	bench.StopTimer()
	rand.Seed(time.Now().Unix())
	n := 256
	A := randomMatrix(n, n)
	B := randomMatrix(n, n)
	bench.StartTimer()
	for i := 0; i < bench.N; i++ {
		Minus(A, B)
    }
}

func BenchmarkSubBLAS_______32(bench *testing.B) {
	bench.StopTimer()
	rand.Seed(time.Now().Unix())
	n := 32
	A := randomMatrix(n, n)
	B := randomMatrix(n, n)
	bench.StartTimer()
	for i := 0; i < bench.N; i++ {
		A.SubBLAS(B)
    }
}

func BenchmarkSub___________32(bench *testing.B) {
	bench.StopTimer()
	rand.Seed(time.Now().Unix())
	n := 32
	A := randomMatrix(n, n)
	B := randomMatrix(n, n)
	bench.StartTimer()
	for i := 0; i < bench.N; i++ {
		A.Sub(B)
    }
}

func BenchmarkMinusBLAS______32(bench *testing.B) {
	bench.StopTimer()
	rand.Seed(time.Now().Unix())
	n := 32
	A := randomMatrix(n, n)
	B := randomMatrix(n, n)
	bench.StartTimer()
	for i := 0; i < bench.N; i++ {
		MinusBLAS(A, B)
    }
}

func BenchmarkMinus__________32(bench *testing.B) {
	bench.StopTimer()
	rand.Seed(time.Now().Unix())
	n := 32
	A := randomMatrix(n, n)
	B := randomMatrix(n, n)
	bench.StartTimer()
	for i := 0; i < bench.N; i++ {
		Minus(A, B)
    }
}

func BenchmarkSubBLASSubM___32(bench *testing.B) {
	bench.StopTimer()
	rand.Seed(time.Now().Unix())
	n := 64
	A := randomMatrix(n, n)
	B := randomMatrix(n, n)
	A = A.SubMatrix(32, 32, 32, 32)
	B = B.SubMatrix(0, 0, 32, 32)
	bench.StartTimer()
	for i := 0; i < bench.N; i++ {
		A.SubBLAS(B)
    }
}

func BenchmarkSubSubM_______32(bench *testing.B) {
	bench.StopTimer()
	rand.Seed(time.Now().Unix())
	n := 64
	A := randomMatrix(n, n)
	B := randomMatrix(n, n)
	A = A.SubMatrix(16, 16, 32, 32)
	B = B.SubMatrix(0, 0, 32, 32)
	bench.StartTimer()
	for i := 0; i < bench.N; i++ {
		A.Sub(B)
    }
}

func BenchmarkMinusBLASSubM__32(bench *testing.B) {
	bench.StopTimer()
	rand.Seed(time.Now().Unix())
	n := 64
	A := randomMatrix(n, n)
	B := randomMatrix(n, n)
	A = A.SubMatrix(16, 16, 32, 32)
	B = B.SubMatrix(0, 0, 32, 32)
	bench.StartTimer()
	for i := 0; i < bench.N; i++ {
		MinusBLAS(A, B)
    }
}

func BenchmarkMinusSubM______32(bench *testing.B) {
	bench.StopTimer()
	rand.Seed(time.Now().Unix())
	n := 64
	A := randomMatrix(n, n)
	B := randomMatrix(n, n)
	A = A.SubMatrix(16, 16, 32, 32)
	B = B.SubMatrix(0, 0, 32, 32)
	bench.StartTimer()
	for i := 0; i < bench.N; i++ {
		Minus(A, B)
    }
}
