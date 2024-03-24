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

# 4. docker設定用のディレクトリを作成
mkdir docker && mkdir docker/php docker/nginx docker/mysql

# 5. 各種Dockerfileと設定ファイルを作成
touch docker/php/Dockerfile docker/nginx/Dockerfile docker/mysql/Dockerfile
touch docker/php/php.ini docker/nginx/default.conf

# 6. 各種Dockerfileと設定ファイルを記述（中身はプロジェクトのファイルを参照）

# 7. docker-compose.ymlの作成（中身はプロジェクトのファイルを参照）

# 8. dockerコンテナの起動（中身はプロジェクトのファイルを参照）
docker-compose up -d
```