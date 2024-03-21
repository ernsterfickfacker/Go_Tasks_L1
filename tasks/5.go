package tasks

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*Разработать программу, которая будет последовательно отправлять
значения в канал, а с другой стороны канала — читать.
По истечению N секунд программа должна завершаться.*/

// отправка сообщений в канал
func senddata(ch chan int, ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		// закрытие канала при закрытии контекста
		case <-ctx.Done():
			fmt.Println("message sending stopped")
			// закрытие канала
			close(ch)
			return
			// иначе продолжаем отправку данных
		default:
			i := rand.Int()
			ch <- i
			fmt.Println("send ", i)
			continue
		}
	}
}

// чтение из канала
func readdata(wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	// сообщения читаются пока канал открыт
	for data := range ch {
		fmt.Println("got data: ", data)
	}
	// при закрытии канала получаем сообщение
	fmt.Println("reading have been stopped")
}

func L1_5() {
	const N = 5
	// канал передачи сообщений
	mainChanel := make(chan int)
	ctx, cansel := context.WithCancel(context.Background())
	wg := sync.WaitGroup{}
	// запускаем потоки записи и чтения параллельно
	wg.Add(2)
	go senddata(mainChanel, ctx, &wg)
	go readdata(&wg, mainChanel)
	// создание таймера
	timer := time.NewTimer(N * time.Second)
	// таймер заблокирован пока не пройдет заданное время и не придет сообщение о том что таймер сработал
	<-timer.C
	cansel()
	// ожидание окончания набора рутин
	wg.Wait()
}
