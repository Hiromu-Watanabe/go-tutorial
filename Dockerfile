# 2020/10/14最新versionを取得
FROM golang:1.20.5-alpine3.18
# アップデートとgitのインストール！！
RUN apk update && apk add git
# appディレクトリの作成
RUN mkdir /go/src/app
# ワーキングディレクトリの設定
WORKDIR /go/src/app
# ホストのファイルをコンテナの作業ディレクトリに移行
ADD . /go/src/app

# go.modファイルとgo.sumファイルの作成
RUN go mod init github.com/Hiromu-Watanabe/go-tutorial
RUN go mod tidy

# ホットリロード（右記参考 : https://zenn.dev/h_sakano/articles/b38336d99f43e4e9e90b）
# RUN go get github.com/oxequa/realize 
# CMD ["realize", "start"]