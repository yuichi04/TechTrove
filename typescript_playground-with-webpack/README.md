# インストール

```shell
# webpack関連
npm install --save-dev webpack webpack-cli webpack-dev-server

# ts関連
npm install --save-dev ts-loader
```

# webpack.config.js の作成

```js
module.exports = {
  entry: {
    bundle: "./src/index.ts",
  },
  output: {
    path: `${__dirname}/dist`,
    filename: "bundle.js",
  },
  mode: "development",
  resolve: {
    extensions: [".ts", ".js"], // from "./index.ts"や"./index.js"を自動認識してくれる
  },
  devServer: {
    static: {
      directory: `${__dirname}/dist`,
    },
    open: true,
  },
  // ファイルに対するルールの設定
  module: {
    rules: [
      {
        test: /\.ts$/, // 拡張子が.tsのファイルをコンパルする
        loader: "ts-loader", // コンパイルにはts-loaderを使用する
      },
    ],
  },
};
```

# コンパイル

```shell
# コンパイル
webpack

# ローカルサーバー起動
webpack serve
```

スクリプトを追加する場合
package.json

```json
"script": {
  "start": "webpack serve",
  "build": "webpack"
}
```
