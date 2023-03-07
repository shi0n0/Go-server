package main

import (
	"html/template"
	"net/http"
	"math/rand"
	"time"
)


func main() {
	// テンプレートをパースする
	tmpl := template.Must(template.ParseFiles("index.html"))

	// HTTPハンドラーを定義する
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// ランダムな数字を生成する
		rand.Seed(time.Now().UnixNano())
		number := rand.Intn(101)

		// HTMLテンプレートにランダムな数字を渡す
		if err := tmpl.Execute(w, number); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	// Webサーバーを開始する
	http.ListenAndServe(":8080", nil)
}

