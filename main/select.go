// Package main provides ...
package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int)

	go setChan(ch)
	for {

		select {
		case <-ch:

			fmt.Println("我收到了something")
		default:
			fmt.Println("我什么也没有收到")
		}
		time.Sleep(1 * time.Second)
	}
}

func setChan(ch chan int) {
	time.Sleep(10 * time.Second)
	ch <- 10
}
