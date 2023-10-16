package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func main() {
	/* POSTに送信するデータを生成 */
	vs := url.Values{}
	vs.Add("id", "1")
	vs.Add("message", "メッセージ")
	fmt.Println(vs.Encode())

	/* POSTメソッドを実行 */
	// res, err := http.PostForm("https://example.com/comments/post", vs)
	_, err := http.PostForm("https://example.com/comments/post", vs)
	if err != nil {
		log.Fatal(err)
	}

}
