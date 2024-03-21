package tasks

import "fmt"

/* Реализовать пересечение двух неупорядоченных множеств. */

func intersection(x []int, y []int) []int {
	var result []int
	cheker := make(map[int]int)
	for _, it := range x {
		cheker[it] = 1
	}
	for _, it := range y {
		_, ok := cheker[it]
		if ok {
			result = append(result, it)
		}
	}
	return result
}

func L1_11() {
	a := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89}
	b := []int{0, 2, 4, 8, 10, 12, 14}
	fmt.Println(intersection(a, b))
}
