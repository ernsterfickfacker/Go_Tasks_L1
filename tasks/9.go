package tasks

import (
	"fmt"
	"sync"
)

/*Разработать конвейер чисел.
Даны два канала: в первый пишутся числа (x) из массива, во второй —
результат операции x*2,
после чего данные из второго канала должны выводиться в stdout.*/

func L1_9() {
	nums := []int{-1, 2, -3, 4, -5, 6, -7, -8, 9, 10}
	wg := sync.WaitGroup{}
	wg.Add(2)
	ch_1 := make(chan int)
	ch_2 := make(chan int)
	// работа первого канала
	go func() {
		defer wg.Done()
		for {
			item, ok := <-ch_1
			if ok {
				// оправка во второй канал
				ch_2 <- item * 2
			} else {
				fmt.Println("Chanel 1 closed")
				break
			}
		}
		// если закрывается первый канал -> прерывается цикл -> закрываем второй канал
		close(ch_2)
	}()
	// работа второго канала
	go func() {
		defer wg.Done()
		for {
			// получение и печать полученных данных
			item, ok := <-ch_2
			if ok {
				fmt.Println(item)
			} else {
				fmt.Println("Chanel 2  closed")
				break
			}
		}
	}()
	// отправка данных в первый канал
	for _, item := range nums {
		ch_1 <- item
	}
	// закрываем канал, когда все данные отправлены
	close(ch_1)
	wg.Wait()
}
