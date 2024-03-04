## Package.json の作成

```shell
npm install -y
```

## インストール

```shell
# 本体
npm install --save-dev webpack

# CLI
npm install --save-dev webpack-cli
```

## バンドルの実行

```shell
# 以下のコマンドを実行すると、ルートディレクトリにdistフォルダが生成され、**その中にバンドルしたファイルが配置される**
npx webpack
```

```json
// package.jsonにスクリプトを追加
"scripts": {
  "build": "webpack"
}
```

## 設定ファイルの作成と書き方

```shell
# 設定ファイルを作成する
touch webpack.config.js
```

```js
module.exports = {
  // エントリーポイントの設定
  entry: "./src/index.js",

  // ファイルの出力設定
  output: {
    // 出力先のパスを決定
    path: `${__dirname}/dist`, // __dirname 現在の絶対パスを取得できる環境変数
    // 出力ファイル名
    filename: "bundle.js",
  },
  mode: "development",
};
```

## webpack のローカルサーバー構築

```shell
npm install --save-dev webpack-dev-server
```

```json
// package.jsonにスクリプトを追加
"start": "webpack server"
// 差分だけビルドする場合
"watch": "webpack --watch"
```

webpack.config.js に追記

```js
devServer: {
  static: "dist", // どのファイルを起動するか
  open: true, // 自動でブラウザを開くか
},
```
