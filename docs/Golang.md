# 環境構築

## 本体のインストール

```shell
brew install go
```

## ディレクトリの作成

```shell
mkdir myproject
cd myproject
```

## モジュールの初期化

```shell
go mod init myproject
```

## サーバーの起動

<u>例）HTTP サーバーを起動し、ブラウザに「Hello, World!」を出力する</u>

### 1. プログラムの作成

#### ファイルを作成

```shell
touch main.go
```

#### ファイルの内容

```go
package main

import (
    "fmt"
    "net/http"
)
// ハンドラ関数を定義
func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World!")
}

func main() {
    // ハンドラ関数をルートパスに割り当て
    http.HandleFunc("/", helloHandler)
    // サーバーをポート8080で起動
    fmt.Println("Starting server at port 8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Println(err)
    }
}
```

### 2. ソースコードのコンパイルとプログラムの実行

```shell
go run main.go
```

<br>
<br>
<br>

# Golang の基本文法

## 1. 変数

データを保存し、プログラム内で操作するための基本的な方法です。

```go
package main

import "fmt"

func main() {
    var a int = 10 // 明示的な型指定
    var b = 20     // 型推論
    c := 30        // 短縮変数宣言

    fmt.Println(a, b, c)
}

// まとめて宣言する場合
var (
    a int
    b string
    c float64
)
```

## 2. 定数

物理定数や設定値など、変更されない固定値を定義するために使用します。

```go
package main

import "fmt"

const Pi = 3.14

func main() {
    fmt.Println(Pi)
}
```

## 3. 基本的なデータ型

様々な種類のデータを表現するために使用します。<br>
Go の基本的なデータ型には、整数、浮動小数点数、文字列、ブール型があります。

```go
package main

import "fmt"

func main() {
    var i int = 42
    var f float64 = 3.14
    var s string = "Hello"
    var b bool = true

    fmt.Println(i, f, s, b)
}
```

## 4. 配列

固定長のデータセットを扱う場合に使用します。

```go
package main

import "fmt"

func main() {
    var arr [3]int = [3]int{1, 2, 3}
    fmt.Println(arr)
}
```

## 5. スライス

リストやダイナミックな配列など、可変長のデータセットを扱う場合に使用します。

```go
package main

import "fmt"

func main() {
    slice := []int{1, 2, 3, 4, 5}
    fmt.Println(slice)
}
```

## 6. マップ

辞書やハッシュマップなど、キーと値のペアを管理するために使用します。

```go
package main

import "fmt"

func main() {
    m := map[string]int{"one": 1, "two": 2}
    fmt.Println(m)
}
```

## 7. 条件分岐

条件に基づいて異なる処理を実行する場合に使用します。

```go
package main

import "fmt"

func main() {
    x := 10
    if x > 5 {
        fmt.Println("x is greater than 5")
    } else {
        fmt.Println("x is less than or equal to 5")
    }
}
```

## 8. ループ

リストの各要素を処理する場合など、同じ処理を繰り返し実行する場合に使用します。

```go
package main

import "fmt"

func main() {
    for i := 0; i < 5; i++ {
        fmt.Println(i)
    }
}
```

## 9. 関数

再利用可能なコードブロックを定義する場合に使用します。例えば、特定の処理を行うコードを関数として定義しておくと便利です。

```go
package main

import "fmt"

func add(a int, b int) int {
    return a + b
}

func main() {
    result := add(2, 3)
    fmt.Println(result)
}
```

## 10. 構造体

ユーザー情報など、複数のデータフィールドを一つのエンティティとして扱う場合に使用します。

```go
package main

import "fmt"

type User struct {
    Name  string
    Email string
}

func main() {
    user := User{Name: "John Doe", Email: "john@example.com"}
    fmt.Println(user)
}
```

## 11. メソッド

ユーザー情報の表示など、構造体に関連する機能を定義する場合に使用します。

```go
package main

import "fmt"

type User struct {
    Name  string
    Email string
}

func (u User) GetDetails() string {
    return fmt.Sprintf("Name: %s, Email: %s", u.Name, u.Email)
}

func main() {
    user := User{Name: "John Doe", Email: "john@example.com"}
    fmt.Println(user.GetDetails())
}
```

## 12. インターフェース

異なる種類のエンティティに共通のメソッドを提供する場合など、異なる型の間で共通の動作を定義する場合に使用します。

```go
package main

import "fmt"

type Describer interface {
    Describe() string
}

type User struct {
    Name  string
    Email string
}

func (u User) Describe() string {
    return fmt.Sprintf("Name: %s, Email: %s", u.Name, u.Email)
}

func main() {
    var d Describer
    user := User{Name: "John Doe", Email: "john@example.com"}
    d = user
    fmt.Println(d.Describe())
}
```

## 13. エラーハンドリング

ファイルの読み込み時にエラーが発生した場合など、エラーが発生した場合にそのエラーを処理するために使用します。

```go
package main

import (
    "errors"
    "fmt"
)

func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

func main() {
    result, err := divide(4, 2)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Result:", result)
    }
}
```

## 14. ゴルーチンとチャンネル

複数のタスクを同時に実行する場合など、並行処理を実現するために使用します。

### ゴルーチン

```go
package main

import (
    "fmt"
    "time"
)

func say(s string) {
    for i := 0; i < 5; i++ {
        time.Sleep(100 * time.Millisecond)
        fmt.Println(s)
    }
}

func main() {
    go say("world")
    say("hello")
}
```

### チャンネルの使用

チャンネルを使ってゴルーチン間でデータをやり取りする場合。

```go
package main

import (
    "fmt"
)

func sum(a []int, c chan int) {
    total := 0
    for _, v := range a {
        total += v
    }
    c <- total
}

func main() {
    a := []int{1, 2, 3, 4, 5}
    c := make(chan int)
    go sum(a, c)
    result := <-c
    fmt.Println(result)
}
```

<br>
<br>
<br>

# その他

## 変数宣言（短縮記法と var の使い分け）について

Golang では、変数宣言には`var`と短縮記法の両方が使用されますが、一般的な使用ケースに応じて使い分けられます。

### 短縮記法（`:=`）

#### 使用シーン

- 関数内で新しい変数を宣言する場合
- 簡潔にコードを書く必要がある場合
- 型推論を利用する場合

#### 例

```go
package main

import "fmt"

func main() {
  x := 10       // int型の変数xを宣言して初期化
  y := "Hello"  // string型の変数yを宣言して初期化
  z := 3.14     // float64型の変数zを宣言して初期化

  fmt.Println(x, y, z)
}
```

### `var`キーワード

#### 使用シーン

- 関数の外で変数を宣言する場合（パッケージレベル）
- 初期値なしで変数を宣言する場合
- 明示的に型を指定する必要がある場合
- 複数の変数を一度に宣言する場合

#### 例

```go
package main

import "fmt"

// パッケージレベルでの変数宣言
var (
  a int
  b string
  c float
)

func main() {
  // 関数内での変数宣言
  var d int = 10
  var e = "World"
  f := 2.718

  // 初期化されていない変数はデフォルト値を持つ
  fmt.Println(a, b, c)
  fmt.Println(d, e, f)
}
```

### 比較

| 特徴             | 短縮記法（`:=`） | `var`キーワード              |
| ---------------- | ---------------- | ---------------------------- |
| 使用場所         | 関数内           | 関数内およびパッケージレベル |
| 型推論           | あり             | あり（初期値がある場合）     |
| 型の明示         | 不要             | 必要に応じて明示             |
| 複数変数の宣言   | 不可             | 可能                         |
| 初期化なしの宣言 | 不可             | 可能                         |

### まとめ

- **短縮記法（`:=`）**:

  - 関数内で新しい変数を宣言する際に一般的に使用されます。
  - コードを簡潔にするために利用され、型推論が行われます。

- **`var`キーワード**:
  - パッケージレベルでの変数宣言や、初期値なしでの宣言、明示的な型指定が必要な場合に使用されます。
  - 複数の変数を一度に宣言する際にも便利です。

実際のコードでは、状況に応じてこれらの方法を使い分けることが重要です。関数内での簡潔な宣言には短縮記法を、明確な型指定やパッケージレベルでの宣言には`var`を使用すると良いでしょう。
