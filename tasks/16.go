package tasks

import "fmt"

/*Реализовать быструю сортировку массива (quicksort) встроенными методами языка.*/

func partition(arr []int, low, high int) ([]int, int) {
	// так как опорный - последний элемент то сравниваем все элементы с ним
	// получим что на позиции i должен стоять опорный элемент а перед ним все элементы меньшие его
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}

func quickSort(arr []int, low, high int) []int {
	if low < high {
		var p int
		// перестановка элементов относительно опорного
		arr, p = partition(arr, low, high)
		// повторить для двух полученных отрезков
		arr = quickSort(arr, low, p-1)
		arr = quickSort(arr, p+1, high)
	}
	return arr
}

func quickSortStart(arr []int) []int {
	// выбор опорного элемента (последний)
	return quickSort(arr, 0, len(arr)-1)
}

func L1_16() {
	somedata := []int{5, 11, 8, 3, 7, 5, 2, 1, 8, 9, 4, 5, 12, 6, 17, 2, 4, 7}
	fmt.Println(quickSortStart(somedata))
}
