package main

import "testing"

func TestDeposit(t *testing.T) {
	if b, err := deposit(100, 50); err != nil || b != 150 {
		t.Errorf("deposit(100, 50) = %g, %v; want 150", b, err)
	}
	if b, err := deposit(100, -5); err == nil || b != 100 {
		t.Error("отрицательное пополнение должно вернуть ошибку и не менять баланс")
	}
}

func TestWithdraw(t *testing.T) {
	if b, err := withdraw(100, 100); err != nil || b != 0 {
		t.Errorf("withdraw(100, 100) = %g, %v; want 0", b, err)
	}
	if b, err := withdraw(100, 150); err == nil || b != 100 {
		t.Error("снятие больше баланса должно вернуть ошибку и не менять баланс")
	}
}
