// Задание 10. Статистика последовательности из N целых чисел.
package main

import "fmt"

type stats struct {
	sum, min, max int
	average       float64
	even, odd     int
}

func calcStats(nums []int) stats {
	s := stats{min: nums[0], max: nums[0]}
	for _, x := range nums {
		s.sum += x
		if x < s.min {
			s.min = x
		}
		if x > s.max {
			s.max = x
		}
		// для отрицательных нечётных x%2 равно -1, поэтому сравниваем с нулём
		if x%2 == 0 {
			s.even++
		} else {
			s.odd++
		}
	}
	s.average = float64(s.sum) / float64(len(nums))
	return s
}

func main() {
	var n int
	fmt.Print("Введите количество чисел N: ")
	if _, err := fmt.Scan(&n); err != nil || n < 1 {
		fmt.Println("Ошибка: N должно быть целым числом не меньше 1")
		return
	}

	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Printf("Число %d: ", i+1)
		if _, err := fmt.Scan(&nums[i]); err != nil {
			fmt.Println("Ошибка: нужно ввести целое число")
			return
		}
	}

	s := calcStats(nums)
	fmt.Println("Сумма:", s.sum)
	fmt.Printf("Среднее: %.2f\n", s.average)
	fmt.Println("Максимальное:", s.max)
	fmt.Println("Минимальное:", s.min)
	fmt.Println("Чётных:", s.even)
	fmt.Println("Нечётных:", s.odd)
}
