package main

import (
	"fmt"
	"log"
	"net/http"

	"go-auth-task-app/routes"
)

func main() {
	// ルーティングの設定
	r := routes.NewRouter()

	// サーバーの起動
	fmt.Println("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
