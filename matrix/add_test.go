// Copyright 2012 Harry de Boer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package matrix

import (
	"math/rand"
	"testing"
	"time"
)

func TestPlus(t *testing.T) {
	A := New(2, 3, []float64{1, 2, 4, 2, 1, 2})
	B := New(2, 3, []float64{8, 4, 2, 2, 1, 0})
	C := Plus(A, B)

	correct := New(2, 3, []float64{9, 6, 6, 4, 2, 2})
	if !equal(C, correct, 0, t) {
		t.FailNow()
	}
}

func TestPlusBLAS(t *testing.T) {
	A := New(2, 3, []float64{1, 2, 4, 2, 1, 2})
	B := New(2, 3, []float64{8, 4, 2, 2, 1, 0})
	C := PlusBLAS(A, B)

	correct := New(2, 3, []float64{9, 6, 6, 4, 2, 2})
	if !equal(C, correct, 0, t) {
		t.FailNow()
	}
}

func TestAdd(t *testing.T) {
	A := New(2, 3, []float64{1, 2, 4, 2, 1, 2})
	B := New(2, 3, []float64{8, 4, 2, 2, 1, 0})
	A.Add(B)

	correct := New(2, 3, []float64{9, 6, 6, 4, 2, 2})
	if !equal(A, correct, 0, t) {
		t.FailNow()
	}
}

func TestAddBLAS(t *testing.T) {
	A := New(2, 3, []float64{1, 2, 4, 2, 1, 2})
	B := New(2, 3, []float64{8, 4, 2, 2, 1, 0})
	A.AddBLAS(B)

	correct := New(2, 3, []float64{9, 6, 6, 4, 2, 2})
	if !equal(A, correct, 0, t) {
		t.FailNow()
	}
}

func TestPlusSubMatrix(t *testing.T) {
	A := New(2, 3, []float64{1, 2, 4, 2, 1, 2})
	B := New(2, 3, []float64{8, 4, 2, 2, 1, 0})
	C := New(2, 2, []float64{6, 6, 2, 2})
	D := A.SubMatrix(0, 1, 2, 2)
	E := B.SubMatrix(0, 1, 2, 2)
	F := Zeros(2, 2)
	F.Plus(D, E)

	if !equal(C, F, 0, t) {
		t.FailNow()
	}
}

func BenchmarkAddBLAS_____1024(bench *testing.B) {
	bench.StopTimer()
	rand.Seed(time.Now().Unix())
	n := 1024
	A := randomMatrix(n, n)
	B := randomMatrix(n, n)
	bench.StartTimer()
	for i := 0; i < bench.N; i++ {
		A.AddBLAS(B)
    }
}

func BenchmarkAdd_________1024(bench *testing.B) {
	bench.StopTimer()
	rand.Seed(time.Now().Unix())
	n := 1024
	A := randomMatrix(n, n)
	B := randomMatrix(n, n)
	bench.StartTimer()
	for i := 0; i < bench.N; i++ {
		A.Add(B)
    }
}

func BenchmarkPlusSBLAS___1024(bench *testing.B) {
	bench.StopTimer()
	rand.Seed(time.Now().Unix())
	n := 1024
	A := randomMatrix(n, n)
	B := randomMatrix(n, n)
	C := Zeros(n, n)
	bench.StartTimer()
	for i := 0; i < bench.N; i++ {
		C.PlusBLAS(A, B)
    }
}

func BenchmarkPlusS_______1024(bench *testing.B) {
	bench.StopTimer()
	rand.Seed(time.Now().Unix())
	n := 1024
	A := randomMatrix(n, n)
	B := randomMatrix(n, n)
	C := Zeros(n, n)
	bench.StartTimer()
	for i := 0; i < bench.N; i++ {
		C.Plus(A, B)
    }
}

func BenchmarkPlusBLAS____1024(bench *testing.B) {
	bench.StopTimer()
	rand.Seed(time.Now().Unix())
	n := 1024
	A := randomMatrix(n, n)
	B := randomMatrix(n, n)
	bench.StartTimer()
	for i := 0; i < bench.N; i++ {
		PlusBLAS(A, B)
    }
}

func BenchmarkPlus________1024(bench *testing.B) {
	bench.StopTimer()
	rand.Seed(time.Now().Unix())
	n := 1024
	A := randomMatrix(n, n)
	B := randomMatrix(n, n)
	bench.StartTimer()
	for i := 0; i < bench.N; i++ {
		Plus(A, B)
    }
}

func BenchmarkAddBLAS______256(bench *testing.B) {
	bench.StopTimer()
	rand.Seed(time.Now().Unix())
	n := 256
	A := randomMatrix(n, n)
	B := randomMatrix(n, n)
	bench.StartTimer()
	for i := 0; i < bench.N; i++ {
		A.AddBLAS(B)
    }
}

func BenchmarkAdd__________256(bench *testing.B) {
	bench.StopTimer()
	rand.Seed(time.Now().Unix())
	n := 256
	A := randomMatrix(n, n)
	B := randomMatrix(n, n)
	bench.StartTimer()
	for i := 0; i < bench.N; i++ {
		A.Add(B)
    }
}

func BenchmarkPlusBLAS_____256(bench *testing.B) {
	bench.StopTimer()
	rand.Seed(time.Now().Unix())
	n := 256
	A := randomMatrix(n, n)
	B := randomMatrix(n, n)
	bench.StartTimer()
	for i := 0; i < bench.N; i++ {
		PlusBLAS(A, B)
    }
}

func BenchmarkPlus_________256(bench *testing.B) {
	bench.StopTimer()
	rand.Seed(time.Now().Unix())
	n := 256
	A := randomMatrix(n, n)
	B := randomMatrix(n, n)
	bench.StartTimer()
	for i := 0; i < bench.N; i++ {
		Plus(A, B)
    }
}

func BenchmarkAddBLAS_______32(bench *testing.B) {
	bench.StopTimer()
	rand.Seed(time.Now().Unix())
	n := 32
	A := randomMatrix(n, n)
	B := randomMatrix(n, n)
	bench.StartTimer()
	for i := 0; i < bench.N; i++ {
		A.AddBLAS(B)
    }
}

func BenchmarkAdd___________32(bench *testing.B) {
	bench.StopTimer()
	rand.Seed(time.Now().Unix())
	n := 32
	A := randomMatrix(n, n)
	B := randomMatrix(n, n)
	bench.StartTimer()
	for i := 0; i < bench.N; i++ {
		A.Add(B)
    }
}

func BenchmarkPlusBLAS______32(bench *testing.B) {
	bench.StopTimer()
	rand.Seed(time.Now().Unix())
	n := 32
	A := randomMatrix(n, n)
	B := randomMatrix(n, n)
	bench.StartTimer()
	for i := 0; i < bench.N; i++ {
		PlusBLAS(A, B)
    }
}

func BenchmarkPlus__________32(bench *testing.B) {
	bench.StopTimer()
	rand.Seed(time.Now().Unix())
	n := 32
	A := randomMatrix(n, n)
	B := randomMatrix(n, n)
	bench.StartTimer()
	for i := 0; i < bench.N; i++ {
		Plus(A, B)
    }
}

func BenchmarkAddBLASSubM___32(bench *testing.B) {
	bench.StopTimer()
	rand.Seed(time.Now().Unix())
	n := 64
	A := randomMatrix(n, n)
	B := randomMatrix(n, n)
	A = A.SubMatrix(32, 32, 32, 32)
	B = B.SubMatrix(0, 0, 32, 32)
	bench.StartTimer()
	for i := 0; i < bench.N; i++ {
		A.AddBLAS(B)
    }
}

func BenchmarkAddSubM_______32(bench *testing.B) {
	bench.StopTimer()
	rand.Seed(time.Now().Unix())
	n := 64
	A := randomMatrix(n, n)
	B := randomMatrix(n, n)
	A = A.SubMatrix(16, 16, 32, 32)
	B = B.SubMatrix(0, 0, 32, 32)
	bench.StartTimer()
	for i := 0; i < bench.N; i++ {
		A.Add(B)
    }
}

func BenchmarkPlusBLASSubM__32(bench *testing.B) {
	bench.StopTimer()
	rand.Seed(time.Now().Unix())
	n := 64
	A := randomMatrix(n, n)
	B := randomMatrix(n, n)
	A = A.SubMatrix(16, 16, 32, 32)
	B = B.SubMatrix(0, 0, 32, 32)
	bench.StartTimer()
	for i := 0; i < bench.N; i++ {
		PlusBLAS(A, B)
    }
}

func BenchmarkPlusSubM______32(bench *testing.B) {
	bench.StopTimer()
	rand.Seed(time.Now().Unix())
	n := 64
	A := randomMatrix(n, n)
	B := randomMatrix(n, n)
	A = A.SubMatrix(16, 16, 32, 32)
	B = B.SubMatrix(0, 0, 32, 32)
	bench.StartTimer()
	for i := 0; i < bench.N; i++ {
		Plus(A, B)
    }
}
