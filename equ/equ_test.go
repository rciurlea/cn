package equ

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleRootBisection(t *testing.T) {
	tcs := []struct {
		f             SVFunc
		a, b, epsilon float64
	}{
		{
			f:       func(x float64) float64 { return x*x - 3 },
			a:       1,
			b:       2,
			epsilon: 0.0001,
		},
	}
	for i, tc := range tcs {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			a1, b1, err := SingleRootBisection(tc.f, tc.a, tc.b, tc.epsilon)
			assert.NoError(t, err)
			assert.True(t, a1 < b1)
			assert.True(t, tc.a <= a1)
			assert.True(t, b1 <= tc.b)
			assert.True(t, b1-a1 <= tc.epsilon)
			assert.InDelta(t, 0, tc.f(a1), 2*tc.epsilon)
			fmt.Println(a1, b1)
		})
	}
}
