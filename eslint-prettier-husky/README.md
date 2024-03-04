# Eslint, Prettier, Husky 設定方法

## package.json の初期化

```shell
npm init -y
```

# ESlint

ESLint：JavaScriptコードの品質を維持し、開発プロセスを改善するのに役立つツール

### インストール

まず、npmを使用してプロジェクトの開発依存関係としてESLintをインストールします。

```shell
npm install eslint --save-dev
```

### ESLintの初期設定

ESLintを初めてプロジェクトに導入する場合は、`eslint --init`コマンドを使用して設定ファイルを生成する。

```shell
npx eslint --init
```

- 設定プロセスでは、プロジェクトのタイプ、使用しているフレームワーク、コードを実行する環境などに関する質問に回答する
- 人気のあるスタイルガイドの適用や、特定のルールのカスタマイズもできる

### ESLintの実行

```shell
npx eslint yourfile.js
```

```shell
npx eslint "src/**/*.js"
```

### スクリプトへの追加

```json
"scripts": {
  "lint": "eslint ."
}
```

### ESLintの無視ファイル

特定のファイルやディレクトリをESLintの検証から除外したい場合は、`.eslintignore`ファイルを作成し、除外したいパスを記述する

```
node_modules
dist
```

# Prettier

Prettier：コードのフォーマットを自動化し、一貫したコーディングスタイルを保つためのツール

### インストール

```shell
npm install --save-dev prettier
```

### 設定ファイルの作成

```shell
touch .prettierrc
```

### 設定ファイルの書き方

```json
{
  "semi": false, // 文末のセミコロンを省略
  "trailingComma": "es5", // ES5で有効な箇所に末尾のカンマを追加
  "singleQuote": true, // シングルクォートを使用
  "printWidth": 80, // 1行の最大文字数を80に設定
  "tabWidth": 2, // タブ幅を2スペースに設定
  "useTabs": false, // スペースを使用してインデント
  "bracketSpacing": true, // オブジェクトリテラルの中括弧にスペースを追加
  "arrowParens": "avoid", // 可能な場合はアロー関数の引数の括弧を省略
  "endOfLine": "lf" // 改行コードをLFに統一
}
```

## Prettierの実行

```shell
npx prettier --write .
```

## スクリプトへの追加

```json
"scripts": {
  "format": "prettier --write ."
}
```

## Prettierの無視ファイル

Prettierから除外したいファイルやディレクトリがある場合は、`.prettierignore`ファイルを作成し、除外したいパスを記述する

```
node_modules
dist
```

以下は、Huskyを使用してGitフックを簡単に管理し、コミット前にLintやテストを自動的に実行する方法について説明するための`README.md`ファイルのサンプルです。このサンプルを基にして、あなたのリポジトリのREADMEを作成し、プロジェクトのニーズに合わせてカスタマイズしてください。

---

以下は、Huskyと`lint-staged`を組み合わせて使用し、ステージングされたファイルに対してLintや自動修正を行う方法について説明するための`README.md`ファイルのサンプルです。このサンプルを基にして、あなたのリポジトリのREADMEを作成し、プロジェクトのニーズに合わせてカスタマイズしてください。

# Husky と lint-staged

このドキュメントでは、Huskyと`lint-staged`をプロジェクトに導入し、Gitコミット時にステージングされたファイルに対して自動的にLintやフォーマットを実行する方法について説明します。これにより、コードの品質を保ちながら、開発プロセスを効率化できます。

### インストール

まず、Huskyと`lint-staged`をプロジェクトの開発依存関係としてインストールします。

```shell
npm install husky lint-staged --save-dev
```

### Huskyの設定

Huskyを設定して、Gitフックを有効にします。

1. **Huskyの有効化**:

   ```shell
   npx husky install
   ```

   `package.json`にHuskyの設定を追加して、`postinstall`スクリプトを設定します。

   ```json
   "scripts": {
     "prepare": "husky install"
   }
   ```

2. **`pre-commit`フックの設定**:

   Huskyを使用して`pre-commit`フックを設定し、`lint-staged`を実行します。

   ```shell
   npx husky add .husky/pre-commit "npx lint-staged"
   ```

### lint-stagedの設定

`lint-staged`を使用して、ステージングされたファイルに対してLintやフォーマットを実行するルールを`package.json`に追加します。

```json
"lint-staged": {
  "*.js": [
    "eslint --fix",
    "prettier --write",
    "git add"
  ]
}
```

この設定では、`.js`ファイルに対して`eslint --fix`と`prettier --write`を実行し、自動修正後にファイルを再ステージングします。

### 使い方

この設定により、`git commit`を実行する際に、自動的に`lint-staged`が呼び出され、設定したルールに基づいてステージングされた`.js`ファイルがLintされ、フォーマットされます。問題がなければコミットが成功し、問題があれば修正を促されます。
