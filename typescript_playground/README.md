# TypeScript 環境構築プロジェクト

## 環境構築手順

### Node.js のインストール

```shell
# 1. 本体のインストール
brew install node

# 2. パスの追加
echo 'export PATH="/usr/local/opt/node@16/bin:$PATH"' >> ~/.zshrc

# 3. コマンドの実行確認
node -v # 16.13.0
```

### TypeScript のインストール

```shell
# 1. 本体のインストール
# グローバルインストールの場合
npm install -g typescript

# プロジェクト固有にインストールする場合
npm install --save-dev typescript

# 2. コマンドの実行確認
#グローバルインストールした場合
tsc -V # Version 4.5.5

# プロジェクト内にインストールした場合
npx tsc -V # Version 4.5.5
```

## tsconfig.json の設定

Node.js はそれ自身では TypeScript をサポートしていないため、TypeScript の導入にあたり TypeScript の設定ファイルである tsconfig.json が必要になる。

### tsconfig.json の生成

```shell
# TypeScriptがpackage.jsonのdependencies(devDependencies)に入っているプロジェクトで実行する
npx tsc --init

# TypeScriptがグローバルインストールされている場合
tsc --init
```

### tsconfig.json の基本的な構成

tsconfig.json ファイルは JSON 形式で記述され、以下のような基本的なセクションで構成される

```
- compilerOptions:
TypeScriptコンパイラの設定を定義する。例えば、出力されるJavaScriptのECMAScriptバージョン（target）、ソースマップの生成（sourceMap）、使用するモジュールシステム（module）などが指定できる。
```

```
- include:
コンパイラがプロジェクトのファイルを解析する際に含めるファイルやディレクトリのパターンの配列。指定しない場合、すべての.ts、.tsx、.d.tsファイルが対象になる（excludeに指定されたファイルを除く）。
```

```
- exclude: コンパイラが無視するファイルやディレクトリのパターンの配列。通常、ビルドプロセスに含めたくないファイル（例：node_modulesディレクトリ）を指定する。
```

```
- files: コンパイルするファイルのリスト。includeやexcludeよりも具体的で、プロジェクト内の特定のファイルだけをコンパイルしたい場合に使用する。
```

#### 基本例

```json
{
  "compilerOptions": {
    "target": "es5",
    "module": "commonjs",
    "strict": true,
    "esModuleInterop": true,
    "outDir": "./dist",
    "sourceMap": true
  },
  "include": ["src/**/*"],
  "exclude": ["node_modules", "dist"]
}

// - target: コンパイルされたJavaScriptのバージョンをES5に設定。
// - module: モジュールシステムとしてCommonJSを使用。
// - strict: すべての厳格な型チェックオプションを有効に。
// - esModuleInterop: CommonJSとESモジュール間の相互運用性を改善。
// - outDir: コンパイルされたファイルを./distディレクトリに出力。
// - sourceMap: ソースマップを生成して、デバッグを容易に。
// - include: srcディレクトリ下のすべてのファイルをコンパイル対象に。
// - exclude: node_modulesとdistディレクトリをコンパイル対象外に。
```

## コンパイラの実行

- TypeScript のコンパイラ（tsc）は、TypeScript のコードを JavaScript に変換するためのコマンドラインツール
- このツールを使用することで、TypeScript で書かれたプログラムやライブラリをブラウザや Node.js などの JavaScript 実行環境で実行可能な形式に変換できる

### 基本的な使用方法

```shell
# TypeScriptを固有のプロジェクトにインストールしている場合は先頭に`npx`をつけること

# 1. 特定のファイルだけコンパイルする場合
# hoge.tsがhoge.jsにコンパイルされる
tsc hoge.ts

# 2. tsconfig.jsonに基づいてコンパイルする場合
tsc

# ウォッチモード
# 開発中にファイルの変更を監視し、変更があった場合に自動的にコンパイルを再実行するには、`tsc`コマンドに`--watch`または`-w`オプションをつける
tsc --watch
```
