package tasks

import (
	"fmt"
	"time"
)

/*Реализовать собственную функцию sleep.*/

func sleep(t time.Duration) {
	to := time.Now().Add(t)
	for time.Now().Unix() < to.Unix() {
	}
}

func L1_25() {
	sleepfor := 10 * time.Second
	start := time.Now().Unix()
	fmt.Println("Start time", start)
	sleep(sleepfor)
	end := time.Now().Unix()
	fmt.Println("End time", end)
	fmt.Printf("Timer was for %d s", end-start)

}
