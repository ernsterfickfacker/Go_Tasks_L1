package tasks

import "fmt"

/*Реализовать бинарный поиск встроенными методами языка.*/
// поиск в отсортированном массиве
// если середине отрезка < искомое значение - продолжаем поиск в правой половине
// иначе в левой
func binarySearch(needle int, haystack []int) bool {
	low := 0
	high := len(haystack) - 1
	for low <= high {
		median := (low + high) / 2
		if haystack[median] < needle {
			low = median + 1
		} else {
			high = median - 1
		}
	}
	if low == len(haystack) || haystack[low] != needle {
		return false
	}
	return true
}

func L1_17() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	number := 10
	res := binarySearch(number, arr)
	if res {
		fmt.Printf("the number %d is contained in the array ", number)
	} else {
		fmt.Printf("the number %d is not contained in the array ", number)

	}

	number1 := 11
	res1 := binarySearch(number1, arr)
	if res1 {
		fmt.Printf("the number %d is contained in the array ", number1)
	} else {
		fmt.Printf("the number %d is not contained in the array ", number1)

	}
}
