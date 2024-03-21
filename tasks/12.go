package tasks

import "fmt"

/*Имеется последовательность строк - (cat, cat, dog, cat, tree)
создать для нее собственное множество. В множестве каждый элемент уникален*/

func L1_12() {
	str := []string{"cat", "cat", "dog", "cat", "tree"}
	set := make(map[string]int)
	var result []string
	for _, it := range str {
		_, ok := set[it]
		if !ok {
			set[it] = 1
			result = append(result, it)
		}
	}
	fmt.Println(result)
}
