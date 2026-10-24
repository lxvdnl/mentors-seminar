// Задание 5. Максимум и минимум из N чисел.
package main

import "fmt"

func minMax(nums []float64) (float64, float64) {
	lo, hi := nums[0], nums[0]
	for _, x := range nums[1:] {
		if x < lo {
			lo = x
		}
		if x > hi {
			hi = x
		}
	}
	return lo, hi
}

func main() {
	var n int
	fmt.Print("Введите количество чисел N: ")
	if _, err := fmt.Scan(&n); err != nil || n < 1 {
		fmt.Println("Ошибка: N должно быть целым числом не меньше 1")
		return
	}

	nums := make([]float64, n)
	for i := 0; i < n; i++ {
		fmt.Printf("Число %d: ", i+1)
		if _, err := fmt.Scan(&nums[i]); err != nil {
			fmt.Println("Ошибка: нужно ввести число")
			return
		}
	}

	lo, hi := minMax(nums)
	fmt.Printf("Максимальное: %g\n", hi)
	fmt.Printf("Минимальное: %g\n", lo)
}
