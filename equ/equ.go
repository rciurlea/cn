package equ

import "errors"

// SVFunc is a single value fuction, e.g. f(x) = x^3 - 2
type SVFunc func(float64) float64

const maxIterations = 100

// SingleRootBisection finds the interval inside which there will be
// a root of f starting from interval [a, b]. f must be continuous
// and must have a single root in the interval. Panics if interval
// is not valid, returns error if iteration limit is exceeded (most
// likely means there are no roots).
func SingleRootBisection(f SVFunc, a, b, epsilon float64) (float64, float64, error) {
	if a >= b {
		panic("invalid interval")
	}
	if epsilon <= 0 {
		panic("invalid epsilon")
	}
	var m float64
	for k := 0; k < maxIterations; k++ {
		m = (a + b) / 2
		if f(m) == 0 {
			return m - epsilon, m + epsilon, nil
		}
		if f(a)*f(m) < 0 {
			b = m
		} else {
			a = m
		}
		if b-a < epsilon {
			return a, b, nil
		}
	}
	return 0, 0, errors.New("max iterations exceeded")
}
