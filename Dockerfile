FROM golang:1.17-alpine

# appディレクトリの作成
RUN mkdir /go/src/Go_project

# copy source
COPY .. /go/src/Go_project/
COPY .. /usr/local/go/src/

# ワーキングディレクトリの設定
WORKDIR /go/src/Go_project

# Environment
ENV LANG C.UTF-8
ENV TZ Asia/Tokyo

# package update
RUN apk update && apk add git

# go.mod パッケージのインストール
RUN go get github.com/cosmtrek/air@v1.29.0
RUN go mod download

# ホットリロードの感知
CMD ["air", "-c", ".air.toml"]