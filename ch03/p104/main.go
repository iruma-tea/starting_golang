package main

import "fmt"

func returnFunc() func() {
	return func() {
		fmt.Println("I'm a fanction")
	}
}

func main() {
	f := returnFunc()
	f()
}
