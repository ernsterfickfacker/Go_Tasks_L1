package tasks

import (
	"fmt"
	"sync"
	"time"
)

/*Написать программу, которая конкурентно рассчитает значение
квадратов чисел взятых из массива (2, 4, 6, 8, 10)
и выведет их квадраты в stdout.*/

func sqrt(i int, ws *sync.WaitGroup) {
	// отложенное выполнение
	// Done - завершение выполнения элемента группы
	defer ws.Done()
	fmt.Println("start ", i)
	// Println выводит в stdout
	// можно показать явно fmt.Fprintln(os.Stdout, ...)
	fmt.Println(i * i)
	// имитация работы
	time.Sleep(2 * time.Second)
	fmt.Println("end ", i)
	//ws.Done()
}

func L1_2() {
	nums := []int{2, 4, 6, 8, 10}
	// определение группы горутин
	ws := sync.WaitGroup{}
	// в которой будет 5 горутин (по числу подсчитываемных элементов)
	ws.Add(5)
	for _, item := range nums {
		go sqrt(item, &ws)
	}
	// ожидание завершения всех горутин в группе
	ws.Wait()
}
