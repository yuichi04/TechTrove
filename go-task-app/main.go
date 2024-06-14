package main

import (
	"fmt"
	"go-task-app/routes"
	"log"
	"net/http"
)

func main() {
	r := routes.NewRouter()
	fmt.Println("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
