package main

import (
	"fmt"
	"time"
)

func enqueue_1(q chan int) {
	for i := 0; i < 100; i++ {
        fmt.Printf("queue info: cap[%d] len[%d]\n", cap(q), len(q))

        fmt.Printf("enqueue:[%d]\n",i)
		q <- i // send i to channel

		if len(q) > 100 {
			time.Sleep(100 * time.Millisecond)
		} else {

            if i%2 == 0 {
			    time.Sleep(1000 * time.Millisecond)
            } else {
			    time.Sleep(2000 * time.Millisecond)
            }
        }
	}
}
func dequeue_1() {
	queue := make(chan int, 10)

	go enqueue_1(queue)

	for i := 0; i < 5; i++ {
		val := <-queue // init val and recv from channel
		fmt.Printf("dequeue:[%d]\n", val)
	}
}

func enqueue_2(q chan int, i int){
    if i%2 == 0 {
        time.Sleep(1000 * time.Millisecond)
    }else{
        time.Sleep(5000 * time.Millisecond)
    }
    q <- i
}
func dequeue_2() {
    queue := make(chan int, 100)

    for i := 0; i < 5; i++ {
        go enqueue_2(queue,i)
    }
	for i := 0; i < 5; i++ {
		val := <-queue // init val and recv from channel
		fmt.Printf("dequeued val:[%d]\n", val)
	}
}

func main() {

    //dequeue_1()

    dequeue_2()
}

