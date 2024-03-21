package tasks

import (
	"fmt"
	"strings"
)

/*Разработать программу, которая проверяет, что все символы в строке уникальные
(true — если уникальные, false etc).
Функция проверки должна быть регистронезависимой.
Например:
abcd — true
abCdefAaf — false
	aabcd — false

*/

func CheckUnique(str string) bool {
	str = strings.ToLower(str)
	m := make(map[rune]int)
	for _, r := range str {
		_, ok := m[r]
		if ok {
			return false
		}
		m[r] = 1
	}
	return true
}

func L1_26() {
	fmt.Println("abcd is unique: ", CheckUnique("abcd"))
	fmt.Println("abCdefAaf is unique: ", CheckUnique("abCdefAaf"))
	fmt.Println("aabcd is unique: ", CheckUnique("aabcd"))
}
