package main

import (
	"fmt"
	"time"
)

func fibonacci(c, quit chan int) {
	time.Sleep(1000 * time.Millisecond)
	x, y := 0, 1
	for {
		select {
		case c <- x:
			fmt.Println("assigning to c")
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		fmt.Println("in go routine")
		for i := 0; i < 10; i++ {
			fmt.Println("in for loop")
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}
