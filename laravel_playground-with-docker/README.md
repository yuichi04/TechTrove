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
