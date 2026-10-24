// Задание 3. Сумма и среднее чисел от 1 до N.
package main

import "fmt"

func sumAndAverage(n int) (int, float64) {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum, float64(sum) / float64(n)
}

func main() {
	var n int
	fmt.Print("Введите N: ")
	if _, err := fmt.Scan(&n); err != nil || n < 1 {
		fmt.Println("Ошибка: N должно быть целым числом не меньше 1")
		return
	}

	sum, avg := sumAndAverage(n)
	fmt.Printf("Сумма от 1 до %d: %d\n", n, sum)
	fmt.Printf("Среднее значение: %.2f\n", avg)
}
