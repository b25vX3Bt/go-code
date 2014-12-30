package main

import (
	"fmt"
	"time"
)

func enqueue(q chan int) {
	for i := 0; i < 100; i++ {
		fmt.Printf("queue cap[%d] len[%d]\n", cap(q), len(q))

		if len(q) > 100 {
			time.Sleep(100 * time.Millisecond)
		}

		q <- i // send i to channel
	}
}

func main() {

	queue := make(chan int, 10)

	go enqueue(queue)

	for i := 0; i < 5; i++ {
		val := <-queue // init val and recv from channel
		fmt.Printf("dequeue:[%d]\n", val)
	}

}
