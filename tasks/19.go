package tasks

import (
	"fmt"
)

/*Разработать программу, которая переворачивает подаваемую
на ход строку (например: «главрыба — абырвалг»). Символы могут быть unicode.*/

func Reverce(s string) string {
	r := []rune(s)
	for i := 0; i*2 < len(r); i++ {
		r[i], r[len(r)-i-1] = r[len(r)-i-1], r[i]
	}
	return string(r)
}

func L1_19() {
	str := "главрыба"
	fmt.Println(Reverce(str))
}
