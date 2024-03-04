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
  // ローカルサーバーの設定
  devServer: {
    static: "dist", // どのファイルを起動するか
    open: true, // 自動でブラウザを開くか
  },
};
