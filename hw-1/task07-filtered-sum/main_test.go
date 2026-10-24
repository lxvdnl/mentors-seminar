package main

import "testing"

func TestFilteredSum(t *testing.T) {
	// подходят 6, 12, 18
	if got := filteredSum(20); got != 36 {
		t.Errorf("filteredSum(20) = %d; want 36", got)
	}
	if got := filteredSum(5); got != 0 {
		t.Errorf("filteredSum(5) = %d; want 0", got)
	}
}
