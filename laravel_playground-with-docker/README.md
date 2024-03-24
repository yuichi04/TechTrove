# PHP, Laravel, Nginx, MySQLの雛形用プロジェクト
## 1. クローンする場合
### 1. リポジトリのクローン
```shell
git clone git@github.com:yuichi04/TechTrove.git
```
### 2. dockerコンテナの起動
```shell
docker-compose up -d
```
## 2. 1から雛形を作成する場合
```shell
# 1. プロジェクトのディレクトリを作成　&　移動
mkdir myproject && cd $_

# 2. プロジェクトのルートディレクトリにlaravelプロジェクトを作成
composer create-project laravel/laravel .

# 3. モジュールのインストール（composer.jsonとpackage.jsonに記述されたモジュール）
composer install && npm install

# 4. Laravel Breezeのインストール
composer require laravel/breeze --dev

# 5. breezeを使うにあたり、必要なファイルを生成（認証機能やTailwindCSSなどが使えるようになる）
php artisan breeze:install

# 6. docker設定用のディレクトリを作成
mkdir docker && mkdir docker/php docker/nginx docker/mysql

# 7. 各種Dockerfileと設定ファイルを作成
touch docker/php/Dockerfile docker/nginx/Dockerfile docker/mysql/Dockerfile
touch docker/php/php.ini docker/nginx/default.conf

# 8. 各種Dockerfileと設定ファイルを記述（中身はプロジェクトのファイルを参照）

# 9. docker-compose.ymlの作成（中身はプロジェクトのファイルを参照）

# 10. dockerコンテナの起動（中身はプロジェクトのファイルを参照）
docker-compose up -d
```
---
## データベース接続設定（環境設定ファイルの修正）
`.env`ファイルの以下の箇所を修正する
```
# Before
DB_CONNECTION=sqlite
# DB_HOST=127.0.0.1
# DB_PORT=3306
# DB_DATABASE=laravel
# DB_USERNAME=root
# DB_PASSWORD=

# After
# docker-compose.ymlに基づいて設定する

DB_CONNECTION=mysql # 接続するデータベース種別
DB_HOST=mysql # service名
DB_PORT=3306 # ports
DB_DATABASE=todoapp # environment > MYSQL_DATABASE
DB_USERNAME=root # environment > MYSQL_ROOT_PASSWORD
DB_PASSWORD=root # デフォルトでは`root`。設定している場合はその値に変更すること
```
---
## マイグレーションの注意点
`php artisan migrate`コマンドは実行する環境の`.env`ファイルに基づいてデータベースに接続するため、ローカル環境ではなく、dockerコンテナ内でコマンドを実行する必要がある。<br>
```shell
docker exec -it [コンテナIDまたは名前] php artisan migrate
```
ローカル環境でコマンドを実行すると以下のようなデータベース接続エラーが発生する。
```shell
 Network address not found: Did you mean to use `sail artisan`? 
      https://laravel.com/docs/sail#executing-artisan-commands

      +41 vendor frames 
```
