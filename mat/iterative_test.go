package mat

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJacobiSolver(t *testing.T) {
	tcs := []struct {
		A, b, x0  *M
		maxIters  uint
		epsilon   float64
		converges bool
	}{
		{
			A:         New(2, 2, 2, 1, 5, 7),
			b:         Vec(11, 13),
			x0:        Vec(1, 1),
			epsilon:   0.0001,
			maxIters:  100,
			converges: true,
		},
		{
			A:         New(3, 3, 2, -1, 1, 2, 2, 2, -1, -1, 2),
			b:         Vec(2, 6, 0),
			x0:        Vec(0, 0, 0),
			epsilon:   0.0001,
			maxIters:  100,
			converges: false,
		},
	}
	for i, tc := range tcs {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			assert.NotPanics(t, func() {
				_, err := SolveJacobi(tc.A, tc.b, tc.x0, tc.epsilon, tc.maxIters)
				if !tc.converges {
					assert.Error(t, err)
					return
				}
				assert.NoError(t, err)
			})
		})
	}
}
