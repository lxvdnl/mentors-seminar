// Задание 9. Мини-банкомат: просмотр, пополнение и снятие с проверкой баланса.
package main

import (
	"errors"
	"fmt"
)

func deposit(balance, amount float64) (float64, error) {
	if amount <= 0 {
		return balance, errors.New("сумма должна быть больше нуля")
	}
	return balance + amount, nil
}

func withdraw(balance, amount float64) (float64, error) {
	if amount <= 0 {
		return balance, errors.New("сумма должна быть больше нуля")
	}
	if amount > balance {
		return balance, errors.New("недостаточно средств на балансе")
	}
	return balance - amount, nil
}

func main() {
	var balance float64
	fmt.Print("Введите начальный баланс: ")
	if _, err := fmt.Scan(&balance); err != nil || balance < 0 {
		fmt.Println("Ошибка: баланс должен быть неотрицательным числом")
		return
	}

	for {
		fmt.Println("\n1. Посмотреть баланс")
		fmt.Println("2. Пополнить баланс")
		fmt.Println("3. Снять средства")
		fmt.Println("0. Выход")
		fmt.Print("Выберите действие: ")

		var choice int
		if _, err := fmt.Scan(&choice); err != nil {
			fmt.Println("Ошибка: нужно ввести номер действия")
			return
		}

		switch choice {
		case 1:
			fmt.Printf("Баланс: %.2f\n", balance)
		case 2, 3:
			var amount float64
			fmt.Print("Введите сумму: ")
			if _, err := fmt.Scan(&amount); err != nil {
				fmt.Println("Ошибка: нужно ввести число")
				continue
			}
			var err error
			if choice == 2 {
				balance, err = deposit(balance, amount)
			} else {
				balance, err = withdraw(balance, amount)
			}
			if err != nil {
				fmt.Println("Ошибка:", err)
			} else {
				fmt.Printf("Готово. Баланс: %.2f\n", balance)
			}
		case 0:
			fmt.Println("До свидания")
			return
		default:
			fmt.Println("Нет такого действия")
		}
	}
}
