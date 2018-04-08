package main

import (
	"fmt"
	"time"
)

func produce(p chan<- int) {
	for i := 0; i < 10; i++ {
		p <- i
		fmt.Println("send:", i)
	}
}

func consumer(c <-chan int) {
	for i := 0; i < 10; i++ {
		v := <-c
		fmt.Println("receive:", v)
	}
}

func main() {
	/*
		ch := make(chan int, 1)
		go func() {
			v := <-ch
			fmt.Println(v)

		}()
		ch <- 1
		fmt.Println("2")
		time.Sleep(1 * time.Second)
	*/
	ch := make(chan int, 5)
	go produce(ch)
	go consumer(ch)
	time.Sleep(1 * time.Second)
}
