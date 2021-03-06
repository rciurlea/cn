package mat

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGaussSimple(t *testing.T) {
	a := New(3, 3, 2, 1, -1, -3, -1, 2, -2, 1, 2)
	b := Vec(8, -11, -3)
	x, err := SolveGaussSimple(a, b)
	assert.NoError(t, err)
	assert.True(t, x.Equals(Vec(2, 3, -1)))
}

func TestGaussPartialPivots(t *testing.T) {
	a := New(3, 3, 2, 1, -1, -3, -1, 2, -2, 1, 2)
	b := Vec(8, -11, -3)
	x, err := SolveGaussPartial(a, b)
	assert.NoError(t, err)
	assert.True(t, x.Equals(Vec(2, 3, -1)))
}
func TestGaussFullPivots(t *testing.T) {
	a := New(3, 3, 2, 1, -1, -3, -1, 2, -2, 1, 2)
	b := Vec(8, -11, -3)
	x, err := SolveGaussFull(a, b)
	assert.NoError(t, err)
	assert.True(t, x.Equals(Vec(2, 3, -1)))
}
