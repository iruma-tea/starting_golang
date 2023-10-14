package main

import (
	"fmt"

	"github.com/iruma-tea/starting_golang/ch03/p120-121/foo"
)

func main() {
	fmt.Println(foo.MAX)
	// fmt.Println(foo.internal_const) // コンパイルエラー
	fmt.Println(foo.FooFunc(5))
	// fmt.Println(foo.internalFunc(5)) // コンパイルエラー
}
