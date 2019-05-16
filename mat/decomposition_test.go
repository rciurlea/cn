package mat

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLUDecomposition(t *testing.T) {
	A := New(3, 3, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	L, U, err := LU(A)
	assert.NoError(t, err)
	assert.True(t, A.Equals(L.Mul(U)))
}

func TestCholeskyDecomposition(t *testing.T) {
	A := New(3, 3, 25, 15, -5, 15, 18, 0, -5, 0, 11)
	B, err := Cholesky(A)
	assert.NoError(t, err)
	assert.True(t, A.Equals(B.Mul(B.Transpose())))
}

func TestColLenght(t *testing.T) {
	tcs := []struct {
		A     *M
		col   int
		panic bool
		norm  float64
	}{
		{
			A:     New(3, 3, 1, 0, 0, 0, 1, 0, 0, 0, 1),
			col:   3,
			panic: false,
			norm:  1,
		},
		{
			A:     New(3, 2, 1, 1, 2, 2, 2, 2),
			col:   1,
			panic: false,
			norm:  3,
		},
		{
			A:     New(3, 3, 1, 0, 0, 0, 1, 0, 0, 0, 1),
			col:   4,
			panic: true,
		},
	}
	for i, tc := range tcs {
		t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
			if !tc.panic {
				var norm float64
				assert.NotPanics(t, func() {
					norm = colLength(tc.A, tc.col)
				})
				assert.True(t, almostEqual(norm, tc.norm))
			} else {
				assert.Panics(t, func() {
					colLength(tc.A, tc.col)
				})
			}
		})
	}
}

func TestGramSchmidt(t *testing.T) {
	A := New(3, 3, 1, 1, 1, -1, 0, 1, 1, 1, 2)
	var B *M
	assert.NotPanics(t, func() {
		B = GramSchmidt(A)
	})
	for i := 1; i <= B.cols; i++ {
		assert.True(t, almostEqual(1, colLength(B, i)))
	}
	var dp float64
	for i := 1; i <= B.cols; i++ {
		for j := 1; j <= B.cols; j++ {
			if i != j {
				dp = 0
				for k := 1; k <= B.rows; k++ {
					dp += B.Get(k, i) * B.Get(k, j)
				}
				assert.Less(t, dp, 0.00000001)
			}
		}
	}
	bsq := B.Mul(B.Transpose())
	for i := 1; i <= bsq.rows; i++ {
		for j := 1; j <= bsq.cols; j++ {
			if i == j {
				assert.True(t, almostEqual(bsq.Get(i, j), 1))
			} else {
				assert.Less(t, bsq.Get(i, j), 0.00000001)
			}
		}
	}
}

func TestQRDecomposition(t *testing.T) {
	A := New(3, 3, 12, -51, 4, 6, 167, -68, -4, 24, -41)
	var Q, R *M
	assert.NotPanics(t, func() {
		Q, R = QR(A)
	})
	assert.True(t, A.Equals(Q.Mul(R)))
}
