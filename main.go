package main

import (
	"fmt"
	"net/http"
)

func oneHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to my website!")
}

func main() {
	http.HandleFunc("/", oneHandler)

	fs := http.FileServer(http.Dir("static/"))                // staticというディレクトリにあるファイルがサーブされる
	http.Handle("/static/", http.StripPrefix("/static/", fs)) // URLのプレフィックスをStripPrefixで除外する

	http.ListenAndServe(":80", nil)
}
