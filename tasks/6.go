package tasks

import (
	"context"
	"fmt"
	"time"
)

// Реализовать все возможные способы остановки выполнения горутины.

// по закрытию канала
func StopRoutine_1() {
	ch := make(chan string)
	go func() {
		for {
			message, err := <-ch
			if !err {
				fmt.Println("End routine")
				return
			}
			fmt.Println(message)
		}
	}()
	ch <- "some message"
	ch <- "another message"
	close(ch)
	time.Sleep(time.Second)
	fmt.Println("End 1")
}

// по закрытию контекста
func StopRoutine_2() {
	ch := make(chan string)
	ctx, cansel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				ch <- "done"
				return
			default:
				fmt.Println("Continue")
				continue
			}
		}
	}(ctx)

	go func() {
		time.Sleep(time.Second)
		cansel()
	}()
	<-ch
	fmt.Println("End 2")
	time.Sleep(2 * time.Second)
}

// с помощью другого канала
func StopRoutine_3() {
	ch := make(chan string)
	ch_1 := make(chan int)
	go func() {
		for true {
			select {
			case ch <- "continue":
				continue
			case <-ch_1:
				ch <- "end"
				close(ch)
				return
			}
		}
	}()

	go func() {
		time.Sleep(1 * time.Second)
		ch_1 <- 0
	}()

	for i := range ch {
		fmt.Println("Got ", i)
	}
	fmt.Println("End 3")
}

func L1_6() {
	StopRoutine_1()
	StopRoutine_2()
	StopRoutine_3()
}
