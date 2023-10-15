package main

import (
	"fmt"

	"github.com/iruma-tea/starting_golang/ch05/p229-230/foo"
)

func main() {
	t := &foo.T{}
	fmt.Println(t.Method1())
	fmt.Println(t.Field1)
	// t.method2() // エラー
	// t.field2　 // エラー
}
