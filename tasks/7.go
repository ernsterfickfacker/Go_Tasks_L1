package tasks

import (
	"fmt"
	"sync"
)

// Реализовать конкурентную запись данных в map.

/* посчитаем суммарное число каждой определенной руны*/

func count(dat []rune, wg *sync.WaitGroup, count map[rune]int, m *sync.Mutex) {
	defer wg.Done()
	for _, item := range dat {
		m.Lock()
		_, ok := count[item]
		if ok {
			count[item] = count[item] + 1
		} else {
			count[item] = 1
		}
		m.Unlock()
	}
}

func L1_7() {
	numbers := make(map[rune]int)
	m := sync.Mutex{}
	wg := sync.WaitGroup{}
	data := []rune{'a', 'b', 'c'}
	data1 := []rune{'b', 'b', 'b', 'c', 'c', 'C', 'd'}
	wg.Add(2)
	go count(data, &wg, numbers, &m)
	go count(data1, &wg, numbers, &m)
	wg.Wait()
	fmt.Println(numbers)
}
