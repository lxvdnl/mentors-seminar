// Задание 6. Количество положительных, отрицательных чисел и нулей.
package main

import "fmt"

func countSigns(nums []float64) (positive, negative, zeros int) {
	for _, x := range nums {
		if x > 0 {
			positive++
		} else if x < 0 {
			negative++
		} else {
			zeros++
		}
	}
	return positive, negative, zeros
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

	positive, negative, zeros := countSigns(nums)
	fmt.Println("Положительных:", positive)
	fmt.Println("Отрицательных:", negative)
	fmt.Println("Нулей:", zeros)
}
