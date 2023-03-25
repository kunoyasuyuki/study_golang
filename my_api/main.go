package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	// helloHandlerという名前で、ハンドラを定義
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		// ハンドラの処理内容:
		// 何がきても、"Hello, World!"の文字列を返す
		io.WriteString(w, "Hello, world!\n")
	}
	// 定義した helloHandler を使うように登録
	http.HandleFunc("/", helloHandler)
	// サーバー起動時のログを出力
	log.Println("server start at port 8080")
	// ListenAndServe関数にて、サーバーを起動
	log.Fatal(http.ListenAndServe(":3000", nil))
}
