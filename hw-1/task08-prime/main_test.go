package main

import "testing"

func TestIsPrime(t *testing.T) {
	primes := []int{2, 3, 5, 13, 97, 7919}
	notPrimes := []int{-7, 0, 1, 4, 9, 25, 100}
	for _, n := range primes {
		if !isPrime(n) {
			t.Errorf("%d должно быть простым", n)
		}
	}
	for _, n := range notPrimes {
		if isPrime(n) {
			t.Errorf("%d не должно быть простым", n)
		}
	}
}
