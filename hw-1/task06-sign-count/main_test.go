package main

import "testing"

func TestCountSigns(t *testing.T) {
	p, n, z := countSigns([]float64{5, -1, 0, 3.5, -0.5, 0})
	if p != 2 || n != 2 || z != 2 {
		t.Errorf("countSigns = %d, %d, %d; want 2, 2, 2", p, n, z)
	}
}
