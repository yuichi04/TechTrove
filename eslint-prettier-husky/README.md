# Eslint, Prettier, Husky 設定方法

## package.json の初期化

```shell
npm init -y
```

## ESLint

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

## Prettier

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

## Husky 設定ガイド

このドキュメントでは、Huskyをプロジェクトに導入し、Gitフックを利用してコードの品質を保持する方法について説明します。Huskyを使用すると、コミット前に自動的にLintやテストを実行するなど、開発プロセスを改善できます。

### インストール

Huskyをプロジェクトの開発依存関係としてnpmを使用してインストールします。

```shell
# husky本体
npm install husky --save-dev

# ステージングされたものに対する操作を行うライブラリ
npm install lint-staged --save-dev
```

### Huskyの設定

Huskyを使用するには、`package.json`にHuskyの設定を追加し、どのGitフックでどのコマンドを実行するかを定義します。

1. **Huskyの有効化**:

   Huskyを有効化するには、以下のコマンドをプロジェクトのルートで実行します。

   ```shell
   npx husky install
   ```

   これにより、`.husky/`ディレクトリがプロジェクトに追加され、Gitフックがこのディレクトリ内で管理されます。

2. **Gitフックの追加**:

   コミット前にLintを実行するための`pre-commit`フックを追加します。

   ```shell
   npx husky add .husky/pre-commit "npm run lint"
   ```

   このコマンドは`.husky/pre-commit`ファイルを作成し、`npm run lint`をコミット前に実行するように設定します。

### Huskyフックのカスタマイズ

Huskyでは、`pre-commit`以外にも多くのフックをカスタマイズできます。例えば、`pre-push`フックを使用して、プッシュ前にテストを実行することも可能です。

```shell
npx husky add .husky/pre-push "npm run test"
```

このコマンドは`.husky/pre-push`ファイルを作成し、`npm run test`をプッシュ前に実行するように設定します。
