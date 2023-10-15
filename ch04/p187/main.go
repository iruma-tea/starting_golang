package main

import "fmt"

// <-chan 受信
// ch:チャネル int 型
func receiver(ch <-chan int) {
	for {
		i := <-ch
		fmt.Println(i)
	}
}

func main() {
	ch := make(chan int)

	go receiver(ch)

	i := 0
	for i < 10000 {
		ch <- i
		i++
	}
}
