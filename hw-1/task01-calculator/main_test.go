package main

import "testing"

func TestCalculate(t *testing.T) {
	cases := []struct {
		a, b float64
		op   string
		want float64
	}{
		{6, 3, "+", 9},
		{6, 3, "-", 3},
		{6, 3, "*", 18},
		{6, 3, "/", 2},
	}
	for _, c := range cases {
		got, err := calculate(c.a, c.b, c.op)
		if err != nil || got != c.want {
			t.Errorf("calculate(%g, %g, %q) = %g, %v; want %g", c.a, c.b, c.op, got, err, c.want)
		}
	}
}

func TestCalculateErrors(t *testing.T) {
	if _, err := calculate(1, 0, "/"); err == nil {
		t.Error("деление на ноль должно вернуть ошибку")
	}
	if _, err := calculate(1, 2, "%"); err == nil {
		t.Error("неизвестная операция должна вернуть ошибку")
	}
}
