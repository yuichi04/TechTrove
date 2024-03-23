FROM php:8.3.2-fpm

# パッケージのインストール
RUN apt-get update && apt-get install -y \
    git \
    zip \
    unzip \
    vim \
    libfreetype6-dev \
    libjpeg62-turbo-dev \
    libmcrypt-dev \
    libpng-dev \
    libfontconfig1 \
    libxrender1 \
    && rm -rf /var/lib/apt/lists/*

# GD, BCMath, PDO_MySQL, MySQLi, Exif拡張をインストール
RUN docker-php-ext-configure gd --with-freetype --with-jpeg \
    && docker-php-ext-install gd bcmath pdo_mysql mysqli exif

# composerをインストール
RUN curl -s http://getcomposer.org/installer | php -- --install-dir=/usr/local/bin --filename=composer

# 使用するNode.jsのバージョンを環境変数として設定
ENV NODE_VERSION=16.16.0

# NVMのインストールとNode.jsのインストール
RUN curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.3/install.sh | bash \
    && . /root/.nvm/nvm.sh \
    && nvm install ${NODE_VERSION} \
    && nvm alias default ${NODE_VERSION} \
    && nvm use default \
    && npm --version \
    && node --version

# Node.jsとnpmのパスを環境変数に追加
ENV PATH="/root/.nvm/versions/node/v${NODE_VERSION}/bin/:${PATH}"

# 作業ディレクトリを/srcに設定
WORKDIR /src

# ホストマシンの現在のディレクトリの内容をコンテナの/src/storageディレクトリに追加
ADD . /src/storage

# /src/storageディレクトリの所有者をwww-dataユーザーとグループに変更
RUN chown -R www-data:www-data /src/storage

# コンテナ起動時に実行するコマンドを設定
ENTRYPOINT [ "bash", "-c", "composer install; npm install; exec php-fpm" ]