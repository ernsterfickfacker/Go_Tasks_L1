package tasks

import "fmt"

/*Удалить i-ый элемент из слайса.*/

func L1_23() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(arr)
	var pos int = 4
	res := append(arr[0:pos], arr[pos+1:]...)
	fmt.Println(res)
}
