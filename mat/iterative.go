package mat

import (
	"fmt"
	"math"
)

// SolveJacobi solves Ax = b using the Jacobi iterative method. Returns
// computed x or an error if the method does not converge. Panics if things
// don't have the correct shape.
func SolveJacobi(A, b, x0 *M, epsilon float64, maxIterations uint) (*M, error) {
	// validate input size
	if A.rows != A.cols {
		panic("square matrix must be provided")
	}
	if b.cols != 1 || b.rows != A.rows {
		panic(fmt.Sprintf("b vector has the wrong shape: %d x %d", b.rows, b.cols))
	}
	if x0.cols != 1 || x0.rows != A.rows {
		panic(fmt.Sprintf("x0 vector has the wrong shape: %d x %d", x0.rows, x0.cols))
	}
	if epsilon < 0 {
		panic("negative error margin")
	}
	xp := x0.Clone() // previous x
	x := xp.Clone()  // start with x same as x0
	var s float64
	for k := 0; k < int(maxIterations); k++ {
		// calculate new x term by term
		for i := 1; i <= x.rows; i++ {
			s = 0
			for j := 1; j <= x.rows; j++ {
				if j != i {
					s += A.Get(i, j) * xp.Get(j, 1)
				}
			}
			x.Set(i, 1, (b.Get(i, 1)-s)/A.Get(i, i))
		}
		// calculate error and stop if appropriate
		s = 0
		for i := 1; i <= x.rows; i++ {
			s += math.Pow((x.Get(i, 1) - xp.Get(i, 1)), 2)
		}
		if math.Sqrt(s) < epsilon {
			return x, nil
		}
		x.CopyTo(xp)
	}
	return nil, fmt.Errorf("iteration limit exceeded")
}
