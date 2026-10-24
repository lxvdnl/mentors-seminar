package main

import "testing"

func TestMinMax(t *testing.T) {
	lo, hi := minMax([]float64{3, -2.5, 10, 0})
	if lo != -2.5 || hi != 10 {
		t.Errorf("minMax = %g, %g; want -2.5, 10", lo, hi)
	}
	lo, hi = minMax([]float64{7})
	if lo != 7 || hi != 7 {
		t.Errorf("minMax одного числа = %g, %g; want 7, 7", lo, hi)
	}
}
