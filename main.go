package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("main execution started")

	//creating go-routines
	go func() {
		fmt.Println("running inside go routine")
	}()

	time.Sleep(10 * time.Millisecond)
	fmt.Println("main execution stopped")
}
