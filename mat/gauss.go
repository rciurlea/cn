package mat

import "fmt"

// SolveUpper solves a linear equation system that has been
// converted to upper triangular form. A is the augmented
// system matrix, i.e. it will have m rows and m+1 columns.
// Panics if A is the wrong size, returns error if the system
// is not properly determined.
func SolveUpper(A *M) (*M, error) {
	if A.cols != A.rows+1 {
		panic("matrix must be of size m/m+1")
	}
	if almostEqual(A.Get(A.rows, A.cols-1), 0) {
		if almostEqual(A.Get(A.rows, A.cols), 0) {
			return nil, fmt.Errorf("undertermined system")
		}
		return nil, fmt.Errorf("overdertermined system")
	}
	x := New(A.rows, 1)
	for i := A.rows; i >= 1; i-- {
		var s float64
		for j := A.cols - 1; j > i; j-- {
			s += A.Get(i, j) * x.Get(j, 1)
		}
		x.Set(i, 1, (A.Get(i, A.cols)-s)/A.Get(i, i))
	}
	return x, nil
}

// GaussSimple tries to bring a matrix to row echelon form with no
// pivoting at all. Can fail if matrix has zeroes on diagonal.
func GaussSimple(a *M) error {
	pivotRow, pivotCol := 1, 1
	for pivotRow <= a.rows && pivotCol <= a.cols {
		pivot := a.Get(pivotRow, pivotCol)
		if almostEqual(pivot, 0) {
			return fmt.Errorf("zero on diagonal")
		}
		// transform all elements below pivot to zeroes
		for i := pivotRow + 1; i <= a.rows; i++ {
			f := a.Get(i, pivotCol) / pivot
			a.Set(i, pivotCol, 0)
			for j := pivotCol + 1; j <= a.cols; j++ {
				a.Set(i, j, a.Get(i, j)-a.Get(pivotRow, j)*f)
			}
		}
		pivotRow++
		pivotCol++
	}
	return nil
}

// GaussPartialPivot brings a matrix to row echelon form by only pivoting
// rows.
func GaussPartialPivot(a *M) {
	pivotRow, pivotCol := 1, 1
	for pivotRow <= a.rows && pivotCol <= a.cols {
		// look for pivots in current column
		rmax, _ := a.MaxIndex(pivotRow, pivotCol, a.rows, pivotCol)
		if almostEqual(a.Get(rmax, pivotCol), 0) {
			// no pivot, move to next column
			pivotCol++
			continue
		}
		a.SwapRows(rmax, pivotRow)
		pivot := a.Get(pivotRow, pivotCol)
		// transform all elements below pivot to zeroes
		for i := pivotRow + 1; i <= a.rows; i++ {
			f := a.Get(i, pivotCol) / pivot
			a.Set(i, pivotCol, 0)
			for j := pivotCol + 1; j <= a.cols; j++ {
				a.Set(i, j, a.Get(i, j)-a.Get(pivotRow, j)*f)
			}
		}
		pivotRow++
		pivotCol++
	}
}

// GaussFullPivot brings a matrix to row echelon form by only pivoting
// rows. Panics if a is not an augmented system matrix. Retuns a permutation
// matrix that can be applied to a x vector to regain initial order.
func GaussFullPivot(a *M) *M {
	if a.cols != a.rows+1 {
		panic("matrix must have shape m/m+1")
	}
	pivotRow, pivotCol := 1, 1
	perm := Eye(a.rows)
	for pivotRow <= a.rows && pivotCol <= a.cols {
		// look for pivots in current column
		rmax, cmax := a.MaxIndex(pivotRow, pivotCol, a.rows, a.cols-1)
		if almostEqual(a.Get(rmax, pivotCol), 0) {
			// no pivot at all, we're done, rest is zeroes
			return perm
		}
		a.SwapRows(rmax, pivotRow)
		a.SwapCols(cmax, pivotCol)
		perm.SwapCols(cmax, pivotCol)
		pivot := a.Get(pivotRow, pivotCol)
		// transform all elements below pivot to zeroes
		for i := pivotRow + 1; i <= a.rows; i++ {
			f := a.Get(i, pivotCol) / pivot
			a.Set(i, pivotCol, 0)
			for j := pivotCol + 1; j <= a.cols; j++ {
				a.Set(i, j, a.Get(i, j)-a.Get(pivotRow, j)*f)
			}
		}
		pivotRow++
		pivotCol++
	}
	return perm
}

// SolveGaussSimple solves the linear equation system  Ax = b
// by gaussian elimination with no pivoting. Panics if A
// is not a square matrix or b is not a vector the size of A.
// Returns error if system does not have a single solution or
// gaussian elimination cannot proceed (zeroes on diagonal).
func SolveGaussSimple(A, b *M) (*M, error) {
	if A.rows != A.cols {
		panic("gaussian solver only works on square matrices")
	}
	if b.rows != A.rows {
		panic("free term vector must be the same size as A")
	}
	a := A.Augment(b)
	err := GaussSimple(a)
	if err != nil {
		return nil, err
	}
	return SolveUpper(a)
}

// SolveGaussPartial solves the linear equation system  Ax = b
// by gaussian elimination with partial pivoting. Panics if A
// is not a square matrix or b is not a vector the size of A.
// Returns error if system does not have a single solution.
func SolveGaussPartial(A, b *M) (*M, error) {
	if A.rows != A.cols {
		panic("gaussian solver only works on square matrices")
	}
	if b.rows != A.rows {
		panic("free term vector must be the same size as A")
	}
	a := A.Augment(b)
	GaussPartialPivot(a)
	return SolveUpper(a)
}

// SolveGaussFull solves the linear equation system  Ax = b
// by gaussian elimination with full pivoting. Panics if A
// is not a square matrix or b is not a vector the size of A.
// Returns error if system does not have a single solution.
func SolveGaussFull(A, b *M) (*M, error) {
	if A.rows != A.cols {
		panic("gaussian solver only works on square matrices")
	}
	if b.rows != A.rows {
		panic("free term vector must be the same size as A")
	}
	a := A.Augment(b)
	perm := GaussFullPivot(a)
	x1, err := SolveUpper(a)
	if err != nil {
		return nil, err
	}
	return perm.Mul(x1), nil
}
