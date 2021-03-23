package main

import (
	"fmt"
	"time"
)

func main() {
	c := producer()
	go consumer(c)

	time.Sleep(1 * time.Millisecond)

}

func producer() chan int {
	out := make(chan int)
	go func() {
		fmt.Println("Hello, I am a Producer")
		for i := 0; i < 20; i++ {
			out <- i * 2
		}
		close(out)
	}()
	return out
}

func consumer(in <-chan int) {
	fmt.Println("Hello I am a consumer")
	for {
		select {
		case v := <-in:
			fmt.Println(v)
		default:
			fmt.Println("Waiting for value from Producer")
		}
	}
}
