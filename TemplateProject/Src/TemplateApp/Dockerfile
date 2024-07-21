#FROMでベースイメージを指定：FROM ubuntu:latest
#RUNでコマンドを実行：RUN apt-get update
#CMDでコンテナ起動時のコマンドを指定：CMD [“echo”, “Hello, Docker!”]
#WORKDIR：Dockerfile内での作業ディレクトリを設定するために使用されます。
# 以降の指令が/appディレクトリで実行されます。
# もしディレクトリが存在しない場合、新しく作成されます。
#COPY：基本的なコピーを提供
#ADD：URLからのダウンロードなど追加の機能を持つ

# goバージョン
FROM golang:1.19.1-alpine
# アップデートとgitのインストール
RUN apk update && apk add git
# boiler-plateディレクトリの作成
#RUN mkdir /go/src/github.com/boiler-plate
RUN mkdir ./go

# URLからのダウンロードなど追加の機能を持つ
#ADD . /go/src/github.com/boiler-plate
# ホストOSの作業ディレクトリをイメージにコピー DockerFileのディレクトリをイメージのServerディレクトリにコピー 
COPY ./ /usr/app/Server

# ワーキングディレクトリの設定　WORKDIRを使用して明示的に作業ディレクトリを設定することが推奨されます。
WORKDIR /usr/app/Server
# 依存関係をインストール　gothubからインポートする手間を省く　go.mod ファイルに定義されたものをインストールする
RUN go mod download
# パッケージのインポート
#RUN go get -u golang.org/x/tools/cmd/goimports

# コンテナ起動時のログ出力
CMD ["echo", "Server Container Start!"]
# main.go 起動
CMD ["go", "run", "main.go"]