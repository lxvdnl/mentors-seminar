// Задание 2. Анализ числа: знак, чётность, делимость на 3 и 5.
package main

import "fmt"

func sign(n int) string {
	if n > 0 {
		return "положительное"
	} else if n < 0 {
		return "отрицательное"
	}
	return "ноль"
}

func isEven(n int) bool {
	return n%2 == 0
}

func divisibleBy(n, d int) bool {
	return n%d == 0
}

func yesNo(ok bool) string {
	if ok {
		return "да"
	}
	return "нет"
}

func main() {
	var n int
	fmt.Print("Введите целое число N: ")
	if _, err := fmt.Scan(&n); err != nil {
		fmt.Println("Ошибка: нужно ввести целое число")
		return
	}

	fmt.Println("Знак:", sign(n))
	if isEven(n) {
		fmt.Println("Чётность: чётное")
	} else {
		fmt.Println("Чётность: нечётное")
	}
	fmt.Println("Делится на 3:", yesNo(divisibleBy(n, 3)))
	fmt.Println("Делится на 5:", yesNo(divisibleBy(n, 5)))
	fmt.Println("Делится и на 3, и на 5:", yesNo(divisibleBy(n, 3) && divisibleBy(n, 5)))
}
