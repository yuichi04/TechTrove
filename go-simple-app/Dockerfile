# ベースイメージとしてGolangの公式イメージを使用
FROM golang:1.22.4

# 作業ディレクトリを作成
WORKDIR /app

# go.modファイルをコンテナの作業ディレクトリにコピーする
COPY go.mod ./

# 依存関係を解決し、必要なパッケージをインストールする
RUN go mod tidy

# ソースコードを作業ディレクトリにコピー
COPY . .

# CompileDaemonをインストールする
RUN go install github.com/githubnemo/CompileDaemon@latest

# CompileDaemonを使用して、ソースコードの変更を監視し、自動的にビルドおよび実行する
# `go-simple-app`の部分は生成されるバイナリファイルの名前を指定するためのもの。任意の名前に変更可能
CMD ["CompileDaemon", "-build=go build -o go-simple-app", "-command=./go-simple-app"]