package main

import (
	"io/ioutil"
	"log"
	"os"
)

func main() {
	f, err := os.Open("foo.txt")
	if err != nil {
		log.Fatal(err)
	}

	/* foo.txtの入力をすべて読み込む */
	bs, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	println(bs)
}
