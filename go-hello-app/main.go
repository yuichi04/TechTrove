package main

import (
	"fmt"
	"log"
	"net/http"
)

// フォームテンプレート
const form = `
  <html>
	  <body>
		  <form action="/greet" method="post">
			  Name: <input type="text" name="name">
				<button type="submit">Greet</button>
		  </form>
	  </body>
	</html>
`

// helloHandlerは、"Hello, World!"と応答するHTTPハンドラです。
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

// greetHandlerは、ユーザーの入力を処理して挨拶を返すHTTPハンドラです。
func greetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		fmt.Fprintf(w, "Hello, %s!", name)
	} else {
		fmt.Fprint(w, form)
	}
}

func main() {
	http.HandleFunc("/", helloHandler)      // ルートURLにhelloHandlerを割り当て
	http.HandleFunc("/greet", greetHandler) // /greet URLにgreetHandlerを割り当て
	fmt.Println("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
