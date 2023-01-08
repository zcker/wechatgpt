FROM golang:1.19-alpine

ENV api_key="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"

ENV wechat_keyword="@#"

RUN export GOPRIVATE=github.com/houko/wechatgpt

WORKDIR /app

COPY . /app

RUN go mod download && go build -o server main.go

CMD ./server