package main

import "testing"

func TestCalcStats(t *testing.T) {
	got := calcStats([]int{4, -3, 7, 0, -8})
	want := stats{sum: 0, min: -8, max: 7, average: 0, even: 3, odd: 2}
	if got != want {
		t.Errorf("calcStats = %+v; want %+v", got, want)
	}
}
