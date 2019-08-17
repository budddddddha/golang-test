# 使用するGolangのイメージを指定する
FROM golang:1.12

ENV GO111MODULE=on

# ワーキングディレクトリを指定する
# WORKDIR /go/src

# RUN ls

# ADD . .

# RUN go get github.com/julienschmidt/httprouter
# RUN go get gopkg.in/gavv/httpexpect.v2

# プロジェクトをビルド
# RUN go build main.go