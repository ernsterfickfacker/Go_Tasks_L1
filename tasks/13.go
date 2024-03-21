package tasks

import "fmt"

/*Поменять местами два числа без создания временной переменной.*/

func L1_13() {
	var x, y int32
	x = 5
	y = 10
	fmt.Println("x = ", x, "y = ", y)
	x, y = y, x
	fmt.Println("x = ", x, "y = ", y)
}
