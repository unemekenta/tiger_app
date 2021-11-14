# ベースとなるDockerイメージ指定
FROM golang:1.15-alpine

# コンテナ内に作業ディレクトリを作成 コンテナログイン時のディレクトリ指定
WORKDIR /go/src/tiger_app

ENV GO111MODULE=on

# パッケージの追加とupdateをしてキャッシュを削除
RUN apk --update add --no-cache \
  alpine-sdk \
  # gooseだけmodで入らなかったので
  && go get bitbucket.org/liamstask/goose/cmd/goose

#ローカルの、go.modとgo.sumをコピー
COPY ./go.mod .
COPY ./go.sum .

# 依存ライブラリをダウンロードする
RUN go mod download

CMD ["/bin/ash"]