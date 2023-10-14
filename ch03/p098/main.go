package main

import "fmt"

func div(a, b int) (int, int) {
	q := a / b // a を bで割った商
	r := a % b // a を bで割った剰余
	return q, r
}

func main() {
	q, r := div(19, 7)
	fmt.Printf("商=%d 剰余=%d\n", q, r)
}
