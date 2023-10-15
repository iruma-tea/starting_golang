package main

import (
	"fmt"

	"github.com/iruma-tea/starting_golang/ch05/p240/foo"
)

func main() {
	t := foo.NewI()
	fmt.Println(t.Method1())
	// fmt.Println(t.method2()) // コンパイルエラー
}
