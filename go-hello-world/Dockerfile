# ベースイメージとしてGolangの公式イメージを使用
FROM golang:1.18

# 作業ディレクトリを作成
WORKDIR /app

# ソースコードを作業ディレクトリにコピー
COPY . .

# アプリケーションをビルド
RUN go build -o hello

# コンテナが起動された時に実行されるコマンドを指定
CMD ["./hello"]