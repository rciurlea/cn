package mat

import "fmt"

// GaussJordan atempts to invert a square matrix A using
// the Gauss-Jordan method. Panics if A is not square and
// returns an error if it's not invertible.
func GaussJordan(A *M) (*M, error) {
	if A.rows != A.cols {
		panic("matrix not square")
	}
	A = A.Augment(Eye(A.rows))
	// create zeroes below diagonal, set diagonal to 1
	for i := 1; i <= A.rows; i++ {
		pivot := A.Get(i, i)
		if almostEqual(pivot, 0) {
			return nil, fmt.Errorf("matrix not invertible")
		}
		// set pivot to 1
		f := 1 / pivot
		A.Set(i, i, 1)
		// scale current row
		for j := i + 1; j <= A.cols; j++ {
			A.Set(i, j, A.Get(i, j)*f)
		}
		// create zeroes under pivot
		for j := i + 1; j <= A.rows; j++ {
			f = A.Get(j, i)
			A.Set(j, i, 0)
			for k := i + 1; k <= A.cols; k++ {
				A.Set(j, k, A.Get(j, k)-f*A.Get(i, k))
			}
		}
	}
	// create zeroes above diagonal
	for i := A.rows; i >= 1; i-- {
		for j := i - 1; j >= 1; j-- {
			f := A.Get(j, i)
			A.Set(j, i, 0)
			for k := i + 1; k <= A.cols; k++ {
				A.Set(j, k, A.Get(j, k)-f*A.Get(i, k))
			}
		}
	}
	return A.Slice(1, A.cols/2+1, A.rows, A.cols), nil
}
