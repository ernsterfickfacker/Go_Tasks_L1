package tasks

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

/*Реализовать постоянную запись данных в канал (главный поток).
Реализовать набор из N воркеров, которые
читают произвольные данные из канала и выводят в stdout.
Необходима возможность выбора количества воркеров при старте.

Программа должна завершаться по нажатию Ctrl+C.
Выбрать и обосновать способ завершения работы всех воркеров.
*/
// посылаем в поток случайные числа
func streamdata(ch chan int) {
	for true {
		ch <- rand.Int()
	}
}

func startworker(i int, wg *sync.WaitGroup, ch chan int, ctx context.Context) {
	defer wg.Done()
	// читаем сообщения из потока
	for work := range ch {
		// сообщаем что работник что-то выполняет
		fmt.Println("worker ", i, " doing ", work)
		// имитация работы
		time.Sleep(1 * time.Second)
		// выход из цикла если получен соотв. сигнал -> контекст закрывается -> ообщение о завершении работы
		select {
		case <-ctx.Done():
			fmt.Println("end of work for ", i)
			return
		default:
			continue
		}
	}
}

func L1_4() {
	// задать число работников
	var workersnum int = 5
	// главный канал
	MainChan := make(chan int)
	// для отлавливания сигналов клавиатуры
	signchan := make(chan os.Signal, 1)
	// искомые типы сигналов
	signal.Notify(signchan, syscall.SIGINT)
	/*SIGHUP - потеря управляющего терминала отправляется
	SIGINT  - прерывание, по умолчанию это ^C (Control-C).
	SIGQUIT - выход, по умолчанию ^\ (Control-Backslash).
	SIGTERM — это общий сигнал для завершения программы.*/
	// контекст чтобы остановка прошла по всей програме
	ctx, cansel := context.WithCancel(context.Background())
	wg := sync.WaitGroup{}
	wg.Add(workersnum)
	// пишем случайные данные в канал
	go streamdata(MainChan)
	// запускаем работников
	for i := 0; i < workersnum; i++ {
		go startworker(i, &wg, MainChan, ctx)
	}
	// ждем SIGINT
	<-signchan
	// завершение
	cansel()
	// ждем окончания всех рутин
	wg.Wait()
}
