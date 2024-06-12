/*
パッケージ宣言

by ChatGPT
- package main: Goのプログラムはパッケージから構成されます。mainパッケージは、実行可能なプログラムを定義する特別なパッケージです。このパッケージにはmain関数が含まれており、プログラムのエントリーポイントとなります。
*/
package main

/*
インポート宣言

by ChatGPT
- import: 他のパッケージをインポートします。
- "fmt": フォーマット済みI/Oを提供するパッケージです。標準出力への出力に使用します。
- "net/http": HTTPクライアントとサーバーの実装を提供するパッケージです。
*/
import (
	"fmt"
	"net/http"
)

/*
ハンドラ関数の定義

by ChatGPT
- func helloHandler(w http.ResponseWriter, r *http.Request): これはHTTPリクエストを処理するためのハンドラ関数です。
  - w http.ResponseWriter: HTTPレスポンスを書き込むためのインターフェースです。
  - r *http.Request: HTTPリクエストの情報を含む構造体です。

- fmt.Fprintf(w, "Hello, World!"): fmt.Fprintfを使って、HTTPレスポンスの本文に「Hello, World!」と書き込みます。wはレスポンスライターで、これを通じてクライアントにデータが送信されます。
*/
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome, Golang!")
}

/*
メイン関数の定義

by ChatGPT
- func main(): プログラムのエントリーポイントであるmain関数を定義します。Goプログラムはこの関数から実行を開始します。
*/
func main() {
	/*
		ルートハンドラの設定

		by ChatGPT
		- http.HandleFunc("/", helloHandler): ルートURL ("/") にアクセスしたときにhelloHandler関数が呼び出されるように設定します。この関数は指定されたパターンのリクエストを処理します。
	*/
	http.HandleFunc("/", helloHandler)

	/*
		サーバー開始メッセージの表示

		by ChatGPT
		- fmt.Println("Starting server at port 8080"): サーバーが起動する際にメッセージを標準出力に表示します。これはデバッグや確認のために役立ちます。
	*/
	fmt.Println("Starng server at port 8080")

	/*
		サーバーの起動

		- if err := http.ListenAndServe(":8080", nil); err != nil {: http.ListenAndServe関数を使ってHTTPサーバーを起動します。この関数は指定されたアドレス（ここではポート8080）でリクエストをリッスンし、処理します。
			- :8080: サーバーがリッスンするポート番号です。
			- nil: デフォルトのマルチプレクサ（http.DefaultServeMux）を使用します。これはhttp.HandleFuncで設定されたハンドラを含みます。

		- if err != nil { fmt.Println(err) }: http.ListenAndServeがエラーを返した場合、そのエラーメッセージを標準出力に表示します。これにより、サーバー起動時のエラーを確認できます。
	*/
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}
}
