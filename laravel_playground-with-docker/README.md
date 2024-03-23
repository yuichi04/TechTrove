# PHP, Laravel, Nginx, MySQLの雛形用プロジェクト
## 1. リポジトリのクローン
```shell
git clone git@github.com:yuichi04/TechTrove.git
```
## 2. dockerコンテナの起動
```shell
docker-compose up -d
```

<br>
<br>
<br>

## 1から雛形を作成する場合の手順
```shell
# 1. プロジェクトのディレクトリを作成　&　移動（myproject部分に任意のプロジェクト名をつける）
mkdir myproject && cd $_

# 2. docker設定用のディレクトリを作成
mkdir docker && mkdir docker/php docker/nginx docker/mysql

# 3. 各種Dockerfileと設定ファイルを作成
touch docker/php/Dockerfile docker/nginx/Dockerfile docker/mysql/Dockerfile
touch docker/php/php.ini docker/nginx/default.conf

# 4. 各種Dockerfileと設定ファイルを記述（中身はプロジェクトのファイルを参照）

# 5. docker-compose.ymlの作成（中身はプロジェクトのファイルを参照）

# 6. プロジェクトのルートディレクトリにlaravelプロジェクトを作成（my-laravel-project部分に任意のlaravelプロジェクト名をつける）
composer create-project --prefer-dist laravel/laravel my-laravel-project

# 7. モジュールのインストール（composer.jsonとpackage.jsonに記述されたモジュール）
composer install && npm install

# 8. dockerコンテナの起動
docker-compose up -d
```