package mat

import (
	"fmt"
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

func TestSlice(t *testing.T) {
	m := New(4, 4, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16)
	n := m.Slice(2, 3, 4, 4)
	assert.True(t, n.Equals(New(3, 2, 7, 8, 11, 12, 15, 16)))
}

func TestEqual(t *testing.T) {
	tcs := []struct {
		A, B *M
		eq   bool
	}{
		{
			A:  New(2, 1, 11, 13),
			B:  New(2, 1, 11, 13),
			eq: true,
		},
	}
	for _, tc := range tcs {
		t.Run(fmt.Sprintf("%s and %s", tc.A, tc.B), func(t *testing.T) {
			assert.True(t, tc.A.Equals(tc.B))
			assert.True(t, tc.B.Equals(tc.A))
		})
	}
}

func TestMul(t *testing.T) {
	m := New(2, 3, 1, 2, 3, 2, 1, 0)
	n := New(3, 4, 1, 0, 1, 2, 2, 3, 1, 0, 1, 1, 1, 2)
	res := m.Mul(n)
	assert.True(t, res.Equals(New(2, 4, 8, 9, 6, 8, 4, 3, 3, 4)))
}

func TestSwapRows(t *testing.T) {
	m := New(3, 4, 1, 0, 1, 2, 2, 3, 1, 0, 1, 1, 1, 2)
	m.SwapRows(1, 3)
	assert.True(t, m.Equals(New(3, 4, 1, 1, 1, 2, 2, 3, 1, 0, 1, 0, 1, 2)))
}

func TestSwapCols(t *testing.T) {
	m := New(2, 3, 1, 2, 3, 2, 1, 0)
	m.SwapCols(1, 3)
	assert.True(t, m.Equals(New(2, 3, 3, 2, 1, 0, 1, 2)))
}

func TestClone(t *testing.T) {
	m := New(2, 3, 1, 2, 3, 2, 1, 0)
	n := m.Clone()
	assert.True(t, m.Equals(n))
	m.Set(1, 1, 20)
	assert.False(t, m.Equals(n))
}

func TestMaxIndex(t *testing.T) {
	m := New(4, 4, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16)
	r, c := m.MaxIndex(2, 2, 4, 4)
	assert.Equal(t, r, 4)
	assert.Equal(t, c, 4)
	r, c = m.MaxIndex(2, 2, 3, 3)
	assert.Equal(t, r, 3)
	assert.Equal(t, c, 3)
}

func TestTranspose(t *testing.T) {
	m := New(2, 3, 1, 3, 5, 2, 4, 6)
	assert.True(t, m.Transpose().Equals(New(3, 2, 1, 2, 3, 4, 5, 6)))
	m = Rand(10, 20)
	assert.True(t, m.Transpose().Transpose().Equals(m))
}

func TestAlmostEqual(t *testing.T) {
	tcs := []struct {
		a, b float64
		eq   bool
	}{
		{1234.567, 1234.567, true},
		{1234.567, 1234.568, false},
		{10000000, 10000000, true},
		{10000000, 10000001, false},
	}
	for _, tc := range tcs {
		t.Run(fmt.Sprintf("%f and %f", tc.a, tc.b), func(t *testing.T) {
			assert.Equal(t, tc.eq, almostEqual(tc.a, tc.b))
		})
	}
}
