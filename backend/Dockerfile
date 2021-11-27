# ベースとなるDockerイメージ指定
FROM golang:1.15-alpine

# コンテナ内に作業ディレクトリを作成 コンテナログイン時のディレクトリ指定
WORKDIR /go/src/tiger_app/backend

ENV GO111MODULE=on

# パッケージの追加とupdateをしてキャッシュを削除
RUN apk --update add --no-cache \
  alpine-sdk

COPY . .

#ローカルのgo.modとgo.sumと.air.tomlをコピー
COPY ./go.mod .
COPY ./go.sum .

COPY ./.air.toml .

# 依存ライブラリをダウンロードし不要なライブラリは削除
RUN go mod tidy && \
    go get -u github.com/cosmtrek/air && \
    go get bitbucket.org/liamstask/goose/cmd/goose

CMD ["air", "-c", ".air.toml"]