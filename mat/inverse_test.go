package mat

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGaussJordan(t *testing.T) {
	m := New(3, 3, 1, 2, 3, 2, 5, 3, 1, 0, 8)
	n, err := GaussJordan(m)
	assert.NoError(t, err)
	assert.True(t, n.Equals(New(3, 3, -40, 16, 9, 13, -5, -3, 5, -2, -1)))
	assert.True(t, m.Mul(n).Equals(Eye(3)))
	assert.True(t, n.Mul(m).Equals(Eye(3)))
}

func TestGaussJordanError(t *testing.T) {
	m := New(3, 3, 1, 2, 3, 4, 5, 6, 2, 4, 6)
	m, err := GaussJordan(m)
	assert.Error(t, err)
}
