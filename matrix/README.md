Package matrix
==============

Package matrix provides matrix multiplication routines.

This package is compatible with Go version 1.


### Goal

I wrote this code as an exercise for me to see how to efficiently implement
matrix multiplication. It is not intended as a package for general usage,
use gomatrix instead: https://code.google.com/p/gomatrix/

Maybe some day part of this code can land up there, but I do not have the
time to look into that currently. Maybe later.


### How to get faster matrix multiplication

I did two (or three) things:

1.	Implement the Strassen (1969) matrix multiplication algorithm, this speeds
	up multiplication of large n x n matrices (for me when n > 80).
	In 1971 Winograd described a variant of this algorithm that uses
	less additions. In 1994 Douglas et al describe a way to place the memory
	in a clever way so you need to allocate less scratch space.
2.	For smaller matrices, use the level 1 BLAS Daxpy function from
	https://github.com/ziutek/blas to speed things up. Note these functions
	only have assembly variants for amd64 bit programs currently.
	The Strassen and Douglas variants fall back to this when they recurse to
	smaller matrices.

Douglas also describes a different way to split the matrix so it should be more
efficient and support non-square matrices. This is not implemented.


### Installation

	go get github.com/harrydb/go/matrix


### Documention

See: https://godoc.org/github.com/harrydb/go/matrix


### Benchmarks

* MulSimple: naive matrix multiplication (but using cache-lines effectively)
* MulGomatrix: the gomatrix implementation.
* MulBLAS: same al MulSimple but the inner loop replaced with the BLAS Daxpy function.
* MulStrassen: the Strassen algorithm
* MulStrassenPar: the Strassen algorithm, but split into two goroutines at each level.
* MulDouglas: Winograd's variant of Strassen's algorithm with Douglas memory placement.

Some uninteresting results where removed

	go test -test.bench . -test.cpu 2,1
	PASS
	BenchmarkMulDouglas__1024	       1	1876168000 ns/op
	BenchmarkMulStrassPar1024-2	       1	1296087000 ns/op
	BenchmarkMulDouglas__512	       5	 264489000 ns/op
	BenchmarkMulStrassPar512-2	      10	 199333100 ns/op
	BenchmarkMulStrassen_512	       5	 298305000 ns/op
	BenchmarkMulBLAS_____512	       5	 562627800 ns/op
	BenchmarkMulSimple___512	       5	 636601200 ns/op
	BenchmarkMulGomatrix_512-2	       5	 783658800 ns/op
	BenchmarkMulGomatrix_512	       5	 659225400 ns/op
	BenchmarkMulDouglas__256	      50	  36059340 ns/op
	BenchmarkMulStrassPar256-2	     100	  34300640 ns/op
	BenchmarkMulStrassen_256	      50	  41638480 ns/op
	BenchmarkMulBLAS_____256	      50	  54319880 ns/op
	BenchmarkMulSimple___256	      20	  75649650 ns/op
	BenchmarkMulGomatrix_256-2	      20	  93713050 ns/op
	BenchmarkMulGomatrix_256	      20	  78700000 ns/op
	BenchmarkMulDouglas__128	     500	   4398502 ns/op
	BenchmarkMulStrassPar128-2	     500	   4666354 ns/op
	BenchmarkMulStrassen_128	     500	   5231164 ns/op
	BenchmarkMulBLAS_____128	     500	   5500720 ns/op
	BenchmarkMulSimple___128	     200	   8934500 ns/op
	BenchmarkMulGomatrix_128-2	     100	  11691120 ns/op
	BenchmarkMulGomatrix_128	     200	   9462230 ns/op
	BenchmarkMulBLAS______64	    5000	    590035 ns/op
	BenchmarkMulSimple____64	    2000	   1013168 ns/op
	BenchmarkMulGomatrix__64	    2000	   1085508 ns/op
	BenchmarkMulBLAS______32	   20000	     95448 ns/op
	BenchmarkMulSimple____32	   10000	    147572 ns/op
	BenchmarkMulGomatrix__32	   10000	    146323 ns/op
