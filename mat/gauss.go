package mat

import "fmt"

// GaussPartialPivot brings a matrix to row echelon form by only pivoting
// rows.
func GaussPartialPivot(a *M) {
	pivotRow, pivotCol := 1, 1
	for pivotRow <= a.rows && pivotCol < a.cols {
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
	if almostEqual(a.Get(a.rows, a.cols-1), 0) {
		if almostEqual(a.Get(a.rows, a.cols), 0) {
			return nil, fmt.Errorf("undertermined system")
		}
		return nil, fmt.Errorf("overdertermined system")
	}
	x := New(a.rows, 1)
	for i := a.rows; i >= 1; i-- {
		var s float64
		for j := a.cols - 1; j > i; j-- {
			s += a.Get(i, j) * x.Get(j, 1)
		}
		x.Set(i, 1, (a.Get(i, a.cols)-s)/a.Get(i, i))
	}
	return x, nil
}
