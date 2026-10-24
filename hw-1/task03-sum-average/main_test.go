package main

import "testing"

func TestSumAndAverage(t *testing.T) {
	cases := []struct {
		n       int
		sum     int
		average float64
	}{
		{1, 1, 1},
		{4, 10, 2.5},
		{10, 55, 5.5},
	}
	for _, c := range cases {
		sum, avg := sumAndAverage(c.n)
		if sum != c.sum || avg != c.average {
			t.Errorf("sumAndAverage(%d) = %d, %g; want %d, %g", c.n, sum, avg, c.sum, c.average)
		}
	}
}
