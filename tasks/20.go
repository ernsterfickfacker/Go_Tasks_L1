package tasks

import (
	"fmt"
	"strings"
)

/*Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow».*/

func ReverceStrWords(s string) string {
	// массив слов (пробел считается сепаратором)
	split := strings.Fields(s)
	// конструктор для строк
	result := strings.Builder{}
	for i := 0; i*2 < len(split); i++ {
		split[i], split[len(split)-i-1] = split[len(split)-1-i], split[i]
	}
	// проходим по словам и записывам в новом порядке
	for _, word := range split {
		result.WriteString(word)
		result.WriteString(" ")
	}
	return result.String()
}

func L1_20() {
	str := "snow dog sun"
	fmt.Println(ReverceStrWords(str))
}
