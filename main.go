package main

import (
	"fmt"
)

func main() {
	fmt.Println("main execution started")
	c := make(chan string)
	//creating go-routines
	go func() {
		fmt.Println("hello " + <-c + " from anonymous goroutine")
	}()

	c <- "Arpeet"
	// time.Sleep(10 * time.Millisecond)
	fmt.Println("main execution stopped")
}
