package main

import "testing"

func TestAnalysis(t *testing.T) {
	if sign(7) != "положительное" || sign(-7) != "отрицательное" || sign(0) != "ноль" {
		t.Error("неверный знак")
	}
	if !isEven(-4) || isEven(-3) || isEven(15) {
		t.Error("неверная чётность")
	}
	if !divisibleBy(15, 3) || !divisibleBy(15, 5) || divisibleBy(10, 3) || !divisibleBy(-30, 5) {
		t.Error("неверная делимость")
	}
}
