# コンテナのビルドと実行

```shell
docker build -t go-simple-app .
docker run --rm -p 8080:8080 -v $(pwd):/app go-simple-app
```

# メモ

## go.mod ファイルの生成

```shell
go mod init go-simple-app
```
