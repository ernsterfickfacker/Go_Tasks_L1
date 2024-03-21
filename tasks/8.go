package tasks

import (
	"fmt"
)

// Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.

func ChangeBit(n int64, p int, v int) int64 {
	if v != 0 {
		// добавить 1 - создаем маску 0000...010....000
		// при побитовой дьзьюнкции это даст 1 в нужной позиции
		mask := int64(1 << (int64(p) - 1))
		fmt.Println(mask)
		return n | mask
	} else {
		// создаем маску обратную случаю для 1
		mask := int64(-1 ^ 1<<(int64(p)-1))
		fmt.Println(mask)
		return n & mask
	}
}

func L1_8() {
	var n int64 = 32
	var pos int = 2
	var v int = 1
	fmt.Println("Number: ", n, "Position", pos, "Bit", v)
	fmt.Println("Result: ", ChangeBit(n, pos, v))
}
