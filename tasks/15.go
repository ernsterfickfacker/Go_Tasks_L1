package tasks

import (
	"fmt"
	"strings"
)

/*К каким негативным последствиям может привести данный фрагмент кода, и как это исправить?
Приведите корректный пример реализации.

Основная проблема заключается в том, что slice - ссылочный тип,
частый вызов функции будет выделять память на строку, котрая будет использоваться лишь частично
*/

/*
var justString string
func someFunc() {
  v := createHugeString(1 << 10)
  justString = v[:100]
}

func main() {
  someFunc()
}
*/

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	//b := make([]byte, 100)
	// никуда не передаем слайс, память можно очистить после окончания работы функции
	//copy(b, v[:100])
	// или через строки
	justString = strings.Clone(v[:100])
	//justString = string(b)
}

func createHugeString(size int) string {
	// встроеннный метод работы со строками для минимизации копирования памяти
	var sb strings.Builder
	for i := 0; i < size; i++ {
		// обавление символа
		sb.WriteString("*")
	}
	return sb.String()
}

func L1_15() {
	someFunc()
	fmt.Println(justString)
}
