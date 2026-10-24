// Задание 7. Сумма чисел от 1 до N, которые чётные и делятся на 3.
package main

import "fmt"

func filteredSum(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		if i%2 == 0 && i%3 == 0 {
			sum += i
		}
	}
	return sum
}

func main() {
	var n int
	fmt.Print("Введите N: ")
	if _, err := fmt.Scan(&n); err != nil || n < 1 {
		fmt.Println("Ошибка: N должно быть целым числом не меньше 1")
		return
	}
	fmt.Printf("Сумма чётных чисел от 1 до %d, которые делятся на 3: %d\n", n, filteredSum(n))
}
