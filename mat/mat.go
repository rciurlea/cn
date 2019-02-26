package mat

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
)

// M is a matrix
type M struct {
	rows int
	cols int
	data []float64
}

// New initializes a matrix of specified size. If initial values
// are provided, these will be used to fill the matrix line by
// line. Only up to rows * cols values will be taken into account.
// Panics if asked to create a 0 size matrix.
func New(rows, cols int, xs ...float64) *M {
	if rows == 0 || cols == 0 {
		panic(fmt.Sprintf("invalid matrix size (%d x %d)", rows, cols))
	}
	m := &M{rows: rows, cols: cols}
	m.data = make([]float64, rows*cols)
	for i, j, k := 1, 1, 0; i <= m.rows && j <= m.cols && k < len(xs); k++ {
		m.Set(i, j, xs[k])
		j++
		if j > m.cols {
			j = 1
			i++
		}
	}
	return m
}

// Set the matrix element at row/col to a new value. Indices are
// 1 based, i.e. upper left corner is row 1 column 1.
// Panics if indices exceed matrix size.
func (m *M) Set(row, col int, value float64) {
	if row > m.rows || row <= 0 || col > m.cols || col <= 0 {
		panic(fmt.Sprintf("invalid matrix indices: %d, %d", row, col))
	}
	m.data[m.rows*(col-1)+row-1] = value
}

// Get value at row/col. Panics if indices exceed matrix size.
func (m *M) Get(row, col int) float64 {
	if row > m.rows || row <= 0 || col > m.cols || col <= 0 {
		panic(fmt.Sprintf("invalid matrix indices: %d, %d", row, col))
	}
	return m.data[m.rows*(col-1)+row-1]
}

// Equals compares matrices for equality. Only makes sense for matrices
// with identical dimensions, panics otherwise.
func (m *M) Equals(other *M) bool {
	if m.rows != other.rows || m.cols != other.cols {
		panic(fmt.Sprintf("trying to compare matrices of different sizes: %dx%d, %dx%d", m.rows, m.cols, other.rows, other.cols))
	}
	for i := 0; i < len(m.data); i++ {
		if !almostEqual(m.data[i], other.data[i]) {
			return false
		}
	}
	return true
}

// Augment matrix with another matrix, returning a new matrix.
// The two matrices must have the same number or rows.
func (m *M) Augment(other *M) *M {
	if m.rows != other.rows {
		panic("can't augment matrices of different heights")
	}
	aug := New(m.rows, m.cols+other.cols)
	mlen := len(m.data)
	for i := 0; i < mlen; i++ {
		aug.data[i] = m.data[i]
	}
	for i := 0; i < len(other.data); i++ {
		aug.data[mlen+i] = other.data[i]
	}
	return aug
}

// Mul multiplies two matrices. If the first has a rows and b columns
// the second must have b rows and c colums. Result has a rows and c cols.
func (m *M) Mul(other *M) *M {
	if m.cols != other.rows {
		panic(fmt.Sprintf("can't multiply matrices of shapes %dx%d and %dx%d", m.rows, m.cols, other.rows, other.cols))
	}
	res := New(m.rows, other.cols)
	var x float64
	for i := 1; i <= res.rows; i++ {
		for j := 1; j <= res.cols; j++ {
			x = 0
			for k := 1; k <= m.cols; k++ {
				x += m.Get(i, k) * other.Get(k, j)
			}
			res.Set(i, j, x)
		}
	}
	return res
}

// SwapRows in place. Panics if row numbers are invalid.
func (m *M) SwapRows(i, j int) {
	if i <= 0 || i > m.rows || j <= 0 || j > m.rows {
		panic(fmt.Sprintf("invalid row numbers: %d %d", i, j))
	}
	var aux float64
	for c := 1; c <= m.cols; c++ {
		aux = m.Get(i, c)
		m.Set(i, c, m.Get(j, c))
		m.Set(j, c, aux)
	}
}

// SwapCols in place. Panics if column numbers are invalid.
func (m *M) SwapCols(i, j int) {
	if i <= 0 || i > m.cols || j <= 0 || j > m.cols {
		panic(fmt.Sprintf("invalid column numbers: %d %d", i, j))
	}
	var aux float64
	for r := 1; r <= m.rows; r++ {
		aux = m.Get(r, i)
		m.Set(r, i, m.Get(r, j))
		m.Set(r, j, aux)
	}
}

// MaxIndex returns the row and column of the largest element
// (absolute value) in the matrix within the specified bounds.
func (m *M) MaxIndex(r1, c1, r2, c2 int) (int, int) {
	if r1 > r2 {
		r1, r2 = r2, r1
	}
	if c1 > c2 {
		c1, c2 = c2, c1
	}
	max := math.Abs(m.Get(r1, c1))
	r, c := r1, c1
	for i := r1; i <= r2; i++ {
		for j := c1; j <= c2; j++ {
			if math.Abs(m.Get(i, j)) > max {
				r, c = i, j
			}
		}
	}
	return r, c
}

// Transpose matrix returning a new matrix.
func (m *M) Transpose() *M {
	t := New(m.cols, m.rows)
	for i := 1; i <= m.rows; i++ {
		for j := 1; j <= m.cols; j++ {
			t.Set(j, i, m.Get(i, j))
		}
	}
	return t
}

// Clone a matrix, returning an identical matrix which shares no memory.
func (m *M) Clone() *M {
	n := New(m.rows, m.cols)
	for i := 0; i < len(m.data); i++ {
		n.data[i] = m.data[i]
	}
	return n
}

// String makes matrices printable
func (m *M) String() string {
	b := &strings.Builder{}
	fmt.Fprintln(b)
	for i := 1; i <= m.rows; i++ {
		for j := 1; j <= m.cols; j++ {
			fmt.Fprintf(b, "%.4g\t", m.Get(i, j))
		}
		fmt.Fprintln(b)
	}
	fmt.Fprintln(b)
	return b.String()
}

// Eye builds the identity matrix
func Eye(n int) *M {
	m := New(n, n)
	for i := 1; i <= n; i++ {
		m.Set(i, i, 1)
	}
	return m
}

// Rand retuns a new matrix with random elements between 0 and 1.
func Rand(rows, cols int) *M {
	m := New(rows, cols)
	for i := 0; i < len(m.data); i++ {
		m.data[i] = rand.Float64()
	}
	return m
}

// Vec creates a column vector (n x 1 matrix) given xs.
func Vec(xs ...float64) *M {
	if len(xs) == 0 {
		panic("can't create empty vector")
	}
	return New(len(xs), 1, xs...)
}

const epsilon = 0.0000001

func almostEqual(a, b float64) bool {
	if a == b {
		return true
	}
	return math.Abs(a-b)/math.Abs(a) < epsilon
}
