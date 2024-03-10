# React + TypeScript + Vite 環境構築手順

## 1. プロジェクトの作成

```shell
npm create vite@latest プロジェクト名
# 現在のディレクトリにインストールする場合はプロジェクト名は`.`とする
```

## 2. フレームワークの選択

```shell
Select a framework: >> - Use arrow-keys. Return to submit.
  Vanilla
  Vue
> React
  Lit
  Svelte
  Solid
  Qwik
  Others
```

## 3. バージョンの選択

```shell
Select a variant: >> - Use arrow-keys. Return to submit.
  TypeScript
> TypeScript + SWC
  JavaScript
  JavaScript + SWC
```

> SWC：Rust ベースの Web プラットフォーム<br>
>
> [SWC 公式](https://swc.rs/)

## 4. パッケージのインストール

```shell
npm install
```

## Import alias の設定

### tsconfig.json

```json
{
  "compilerOptions": {
    "target": "ES2020",
    "useDefineForClassFields": true,
    "lib": ["ES2020", "DOM", "DOM.Iterable"],
    "module": "ESNext",
    "skipLibCheck": true,

    /* Bundler mode */
    "moduleResolution": "bundler",
    "allowImportingTsExtensions": true,
    "resolveJsonModule": true,
    "isolatedModules": true,
    "noEmit": true,
    "jsx": "react-jsx",

    /* Linting */
    "strict": true,
    "noUnusedLocals": true,
    "noUnusedParameters": true,
    "noFallthroughCasesInSwitch": true,

    <!-- 以下を追加する -->
    /* Import alias */
    "baseUrl": "./",
    "paths": {
      "@/*": ["src/*"]
    }
  },
  "include": ["src"],
  "references": [{ "path": "./tsconfig.node.json" }]
}
```

> **Import alias**：
> インポートを以下のように簡単に記述できるように設定するもの
> `'../../../../components/button'`→ `'@/components/button'`

### vite.config.ts

1. モジュールのインストール

```shell
npm install --save-dev vite-tsconfig-paths
```

2. モジュールの使用

```typescript
import { defineConfig } from "vite";
import react from "@vitejs/plugin-react-swc";
import tsconfigPaths from "vite-tsconfig-paths"; // インポートを追加

export default defineConfig({
  plugins: [react(), tsconfigPaths()], // tsconfigPaths()を追加
});
```

## 5. テストの設定

### モジュールのインストール

```shell
npm install --save-dev vitest happy-dom @vitest/coverage-v8 @testing-library/react @testing-library/user-event @testing-library/jest-dom
```

> happy-dom：js-dom よりも高速（テスト環境で有効）

#### テストの実行

`package.json`にテスト実行用のスクリプトを追加

```json
"scripts": {
  "test": "vitest",
  "test:watch": "vitest run --coverage",
  "coverage": "vitest run --coverage"
}
```

### vitest の設定

> ※jest-dom のインポートをテストファイルごとに記述するのが面倒な場合は設定する
> 設定することでインポート文を記述しなくても自動で読み込んでくれるようになる

1. プロジェクトのルートディレクトリに設定ファイルを作成

```
touch vitest-setup.ts
```

2. 以下を設定ファイルに記述

> vitest-setup.ts

```typescript
import "@testing-library/jest-dom/vitest";
```

> vite.config.ts

```typescript
/// <reference types="vitest" />
import { defineConfig } from "vite";
import react from "@vitejs/plugin-react-swc";
import tsconfigPaths from "vite-tsconfig-paths";

export default defineConfig({
  plugins: [react(), tsconfigPaths()],
  test: {
    globals: true,
    environment: "happy-dom",
    setupFiles: ["./vitest-setup.ts"],
  },
});
```

> tsconfig.json

```json
"include": ["src", "vitest-setup.ts"],
```

## 6. ESLint の設定

### パッケージのインストール

```shell
npm install --save-dev eslint
```

### 初期化

```shell
npx eslint --init
```

```shell
? How would you like to use ESLint?
  To check syntax only
❯ To check syntax and find problems
  To check syntax, find problems, and envorce code style

? How would you like to use ESLint?
❯ JavaScript modules (import/export)
  CommonJS (require/exports)
  None of these

? Witch framework does your project use?
❯ React
  Vue.js
  None of these

? Dose your project use TypeScript?
  No
❯ Yes

? Where does your code run? (Press <space> to select, <a> to toggle all, <i> to invert selection)
❯ Browser
  Node

? What format do you want your config file to be in?
❯ JavaScript
  YAML
  JSON

? Would you like to install them now?
  No
❯ Yes

? Which package manager do you want to use?
❯ npm
  yarn
  pnpm
```

### ESLint 実行用のスクリプトを追加

```json
"scripts": {
    ...
    "lint": "eslint src",
    ...
  },
```

### React の未使用の場合はインポートをしなくてもエラーが出ないように設定

> .eslintrc.cjs

```cjs
rules: {
    "react/react-in-jsx-scope": "off",
  },
```

これにより、以下のようにインポートしても構文エラーが発生しなくなる

```jsx
import { useState } from "react";

import React, { useState } from "react"; // この記述のReact部分が不要になる
```

### lint チェック時に修正まで行いたい場合の設定

> package.json

```json
"scripts": {
  "lint": "eslint src",
  "lint:fix": "eslint src --fix" // これを追加
}
```

## 7. Prettier の設定

### パッケージのインストール

```shell
npm install --save-dev prettier
```

### 設定ファイルの作成

プロジェクトのルートディレクトリに`prettier.config.js`を作成し、以下を記述

```js
/** @type {import("prettier").Config} */
const config = {
  semi: true,
  tabWidth: 2,
  singleQuote: true,
};

export default config;
```

### Prettier 実行用のスクリプトを追加

> package.json

```json
"scripts": {
  "format": "prettier . --write"
}
```

## 7. Husky と LintStaged の設定

> ※以下は全てプロジェクトのルートディレクトリで行う必要がある

### git の初期化

```shell
git init
```

### パッケージのインストール

```shell
npm install --save-dev husky lint-staged

npx husky install
```

### スクリプトの追加

> package.json

```json
"scripts": {
  "prepare": "husky install"
},
"lint-staged": {
  "*.{js,jsx,ts,tsx}": [
    "prettier --write",
    "eslint --fix"
  ]
}
```

### husky コマンドを追加

```shell
npx husky add .husky/pre-commit "npx lint-staged"
```

## 8. Storybook の設定

### 初期化

```shell
npx storybook init --builder @storybook/builder-vite
```

以下を質問されたら`yes`で OK（.eslintrc.cjs に自動で storybook に関連する便利プラグインを入れてくれる）

```shell
We have detected that you're using ESLint. Storybook provides a plugin that gives the best experience with Storybook and helps follow best practices: https://github.com/storybookjs/eslint-plugin-storybook#readme

Would you like to install it? … yes
```

## 9. TailwindCSS の導入

### パッケージのインストール

```shell
npm install --save-dev tailwindcss postcss autoprefixer
```

### 初期化

```shell
npx tailwindcss init -p
```

`tailwindcss.config.js`と`postcss.config.js`が生成される

### tailwindcss.config.js の設定

> tailwindcss.config.js

```javascript
/** @types {import('tailwindcss').Config} */
exoprt default {
  content: ['./index.html', './src/**/*.{js,ts,jsx,tsx}'],
  theme: {
    extend: {},
  },
  plugins: [],
};
```

### index.css の修正

> index.css

以下を記述する

```css
@tailwind base;
@tailwind components;
@tailwind utilities;
```

## 10. Shadcn の導入

### パッケージのインストール

```shell
npx shadcn-ui@latest init
```

```shell
✔ Would you like to use TypeScript (recommended)? … no / yes
✔ Which style would you like to use? › Default
✔ Which color would you like to use as base color? › Slate
✔ Where is your global CSS file? … src/index.css
✔ Would you like to use CSS variables for colors? … no / yes
✔ Are you using a custom tailwind prefix eg. tw-? (Leave blank if not) …
✔ Where is your tailwind.config.js located? … tailwind.config.js
✔ Configure the import alias for components: … @/components
✔ Configure the import alias for utils: … @/lib/utils
✔ Are you using React Server Components? … no / yes
✔ Write configuration to components.json. Proceed? … yes
```

### tailwind.config.js の修正

> tailwind.config.js

```javascript
// 1. `module.exports`を`export default`に変更
module.exports = {
  darkMode: ["class"],
  // 以下省略
}
export default {
  darkMode: ["class"],
  // 以下省略
}

// 2. `require`を`import`に変更
plugins: [require("tailwindcss-animate")],
plugins: [import("tailwindcss-animate")],
```

## 10. GitHub Actions（CI） の導入

### ルートディレクトリに`.github/workflows`を作成

```shell
mkdir ./github/workflows
```

### yml ファイルを作成

```shell
# ./github/workflows
touch hogehoge.yml

# 例
touch lint_test.yml
```

### yml ファイルの中身を記述する

> lint_test.yml

```yml
name: Lint and Test # GitHub Actionで実行する処理の名前

on: push # トリガーの設定。ここではpush時に設定

jobs:
  lint-and-test:
    runs-on: ubuntu-latest

    steps:
      - uses: actions/checkout@v4

      - name: Setup Node
        uses: actions/setup-node@v4
        with:
          node-version: "20"

      - name: Install Dependencies
        run:
          npm ci
          # `npm install`は`package.json`の中身を見てインストールを実行する
          # `npm ci`は`packagelock.json`の中身を見てインストールを実行する

      - name: Run Lint
        run: npm run lint

      - name: Run Test
        run: npm run test
```

## 11. GitHub Actions（CD）の導入

[公式 - 静的サイトのデプロイ](https://ja.vitejs.dev/guide/static-deploy.html)
