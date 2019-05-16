package mat

import (
	"fmt"
	"math"
)

// LU computes the LU decomposition of square matrix A, returning
// new matrices L and U. A is not mutated. Panics if A is not square.
// Returns error if operation cannot be completed (e.g. singular matrix).
func LU(A *M) (*M, *M, error) {
	if A.cols != A.rows {
		panic("matrix must be square for LU decomposition")
	}
	// handle trivial case (size 1)
	if A.cols < 2 {
		return New(1, 1, A.Get(1, 1)), Eye(1), nil
	}
	U := A.Clone()
	L := Eye(A.rows)
	for k := 1; k <= U.cols; k++ {
		pivot := U.Get(k, k)
		if almostEqual(pivot, 0) {
			if L.Mul(U).Equals(A) {
				return L, U, nil
			}
			return nil, nil, fmt.Errorf("zero pivot found at %d, %d", k, k)
		}
		// transform all elements below pivot to zeroes, saving scaling
		// factors to L matrix
		for i := k + 1; i <= U.rows; i++ {
			f := U.Get(i, k) / pivot
			L.Set(i, k, f)
			U.Set(i, k, 0)
			for j := k + 1; j <= U.cols; j++ {
				U.Set(i, j, U.Get(i, j)-U.Get(k, j)*f)
			}
		}
	}
	return L, U, nil
}

// Cholesky decomposes a symmetrical matrix A into it's B*B' representation.
// Panics if A is not square or symmetrical.
func Cholesky(A *M) (*M, error) {
	if A.rows != A.cols {
		panic("need square matrix for Cholesky decomposition")
	}
	if !A.Equals(A.Transpose()) {
		panic("need symmetrical matrix for Cholesky decomposition")
	}
	B := New(A.rows, A.cols)
	var s float64
	for i := 1; i <= A.rows; i++ {
		// diagonal element
		s = 0
		for k := 1; k < i; k++ {
			s += B.Get(i, k) * B.Get(i, k)
		}
		s = A.Get(i, i) - s
		if s < 0 {
			return nil, fmt.Errorf("negative diagonal value at %d, %d", i, i)
		}
		B.Set(i, i, math.Sqrt(s))
		// other elements
		for j := i + 1; j <= A.rows; j++ {
			s = 0
			for k := 1; k < i; k++ {
				s += B.Get(j, k) * B.Get(i, k)
			}
			s = (A.Get(j, i) - s) / B.Get(i, i)
			B.Set(j, i, s)
		}
	}
	return B, nil
}

// QR computes the QR decomposition of square matrix A, returning
// orthogonal matrix Q and upper triangular matrix R. A is not mutated.
// Panics if A is not square.
func QR(A *M) (*M, *M) {
	Q := GramSchmidt(A)
	R := Q.Transpose().Mul(A)
	return Q, R
}

// GramSchmidt generates a orthonormal base for the column space of A.
// A needs to have linearly independent columns.
func GramSchmidt(A *M) *M {
	// panic if A is "wide", it certainly can't have independent columns
	if A.cols > A.rows {
		panic("A needs to have independent columns")
	}
	B := A.Clone()
	var dp, s float64
	for i := 1; i <= B.cols; i++ {
		// column i in B already contains column i in A, due to cloning
		// move on to projection subtraction
		for j := 1; j < i; j++ {
			// calculate dot product between col i in A and col j in B
			// also calculate squared norm of col j in B
			dp, s = 0, 0
			for k := 1; k <= A.rows; k++ {
				dp += A.Get(k, i) * B.Get(k, j)
				s += B.Get(k, j) * B.Get(k, j)
			}
			// adjust dot product by squared norm
			dp /= s
			// and subtract col j in B scaled by calculated factor
			for k := 1; k <= A.rows; k++ {
				B.Set(k, i, B.Get(k, i)-B.Get(k, j)*dp)
			}
		}
	}
	// normalize vectors in B
	for i := 1; i <= B.cols; i++ {
		norm := colLength(B, i)
		for j := 1; j <= B.rows; j++ {
			B.Set(j, i, B.Get(j, i)/norm)
		}
	}
	return B
}

func colLength(A *M, col int) float64 {
	if col < 1 || col > A.cols {
		panic("invalid column number")
	}
	var s float64
	for i := 1; i <= A.rows; i++ {
		s += A.Get(i, col) * A.Get(i, col)
	}
	return math.Sqrt(s)
}
