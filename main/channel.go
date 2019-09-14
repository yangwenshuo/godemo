package main

import "fmt"

func Count(ch chan int, i int) {
	//这里是会阻塞的,直到有goruntine把channel的东西拿出来才会继续往下跑到下面的打印
	ch <- 1

	fmt.Println("来来" + fmt.Sprintf("%d", i))
}

func main() {
	chs := make([]chan int, 10, 10)
	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
		go Count(chs[i], i)
	}

	for i, ch := range chs {
		//这里也会阻塞的,直到有人往这个chan里赛东西
		o := <-ch

		fmt.Println("indxe" + fmt.Sprintf("%d", i) + "value" + fmt.Sprintf("%d", o))
	}
}
