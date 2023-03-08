package main

import (
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

// ランダムな数字を生成する関数
func RandomNumber() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(101)
}

// テンプレートをレンダリングするハンドラー
func renderTemplate(w http.ResponseWriter, r *http.Request) {
	// テンプレート設定
	tmpl := template.Must(template.ParseFiles("index.html"))

	// ボタンを押されたときにランダムな数字を生成
	if r.Method == http.MethodPost {
		randomNumber := RandomNumber()
		tmpl.Execute(w, randomNumber)
		return
	}

	// 初回アクセス時はランダムな数字を表示しない
	tmpl.Execute(w, "xxx")
}

func main() {
	http.HandleFunc("/", renderTemplate)
	http.ListenAndServe(":8080", nil)
}