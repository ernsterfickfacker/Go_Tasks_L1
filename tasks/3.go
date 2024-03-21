package tasks

import (
	"fmt"
	"sync"
)

/*Дана последовательность чисел: 2,4,6,8,10.
Найти сумму их квадратов(2**2+3**2+4**2….)
с использованием конкурентных вычислений.*/
/*
Почти как предыдущее, но теперь нужно записать результат в одну переменную
Чтобы разграничить доступ к ресурсу используем мьютекс
*/

func sqrt_sum(i int, sum *int, ws *sync.WaitGroup, m *sync.Mutex) {
	defer ws.Done()
	// локируем доступ
	m.Lock()
	*sum = *sum + i*i
	// разблокируем доступ
	m.Unlock()
}

func L1_3() {
	nums := []int{2, 4, 6, 8, 10}
	ws := sync.WaitGroup{}
	m := sync.Mutex{}
	ws.Add(5)
	var s int = 0
	for _, item := range nums {
		go sqrt_sum(item, &s, &ws, &m)
	}
	ws.Wait()
	fmt.Println(s)
}
