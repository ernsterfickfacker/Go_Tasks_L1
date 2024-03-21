package tasks

import (
	"fmt"
	"sync"
)

/*Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
По завершению программа должна выводить итоговое значение счетчика.*/

type Counter struct {
	value int
}

func (c *Counter) Add(m *sync.Mutex) {
	m.Lock()
	c.value += 1
	m.Unlock()
}

func L1_18() {
	count := Counter{
		0,
	}
	wg := sync.WaitGroup{}

	m := sync.Mutex{}
	var n int = 15
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(c *Counter, m *sync.Mutex) {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				c.Add(m)
			}
		}(&count, &m)
	}
	wg.Wait()
	fmt.Println(count.value)
}
