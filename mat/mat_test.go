package mat

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmptyMatrixInitialization(t *testing.T) {
	m := New(3, 4)
	assert.Equal(t, 3, m.rows)
	assert.Equal(t, 4, m.cols)
}

func TestMatrixSet(t *testing.T) {
	tcs := []struct {
		r, c int
		val  float64
		idx  uint
	}{
		{r: 1, c: 1, val: 5, idx: 0},
		{r: 3, c: 1, val: 6, idx: 2},
		{r: 2, c: 3, val: 10, idx: 9},
	}
	m := New(4, 5)
	for _, tc := range tcs {
		m.Set(tc.r, tc.c, tc.val)
		assert.Equal(t, tc.val, m.data[tc.idx], "expected to find %f at index %d", tc.val, tc.idx)
	}
}

func TestMatrixInitWithData(t *testing.T) {
	m := New(3, 3, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	assert.Equal(t, 1.0, m.data[0])
	assert.Equal(t, 4.0, m.data[1])
	assert.Equal(t, 7.0, m.data[2])
	assert.Equal(t, 2.0, m.data[3])
	assert.Equal(t, 5.0, m.data[4])
	assert.Equal(t, 8.0, m.data[5])
	assert.Equal(t, 3.0, m.data[6])
	assert.Equal(t, 6.0, m.data[7])
	assert.Equal(t, 9.0, m.data[8])
}

func TestMatrixEquals(t *testing.T) {
	m := New(2, 3, 1, 2, 3, 4, 5, 6)
	n := New(2, 3, 1, 2, 3, 4, 5, 6)
	p := New(2, 3, 1, 0, 0, 0, 1, 0)
	assert.True(t, m.Equals(n))
	assert.False(t, m.Equals(p))
}

func TestEye(t *testing.T) {
	m := Eye(3)
	assert.True(t, m.Equals(New(3, 3, 1, 0, 0, 0, 1, 0, 0, 0, 1)))
}

func TestAugment(t *testing.T) {
	m := New(2, 2, 1, 2, 3, 4)
	n := New(2, 3, 11, 12, 13, 14, 15, 16)
	assert.True(t, m.Augment(n).Equals(New(2, 5, 1, 2, 11, 12, 13, 3, 4, 14, 15, 16)))
}