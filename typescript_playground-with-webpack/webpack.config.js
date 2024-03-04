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
