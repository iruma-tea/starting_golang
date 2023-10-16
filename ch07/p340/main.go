package main

import (
	"log"
	"net/http"
)

func main() {
	/* GETメソッドでURLにアクセス */
	// res, err := http.Get("https://www.google.com/")
	_, err := http.Get("https://www.google.com/")
	if err != nil {
		log.Fatal(err)
	}
}
