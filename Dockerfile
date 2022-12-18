FROM golang:1.20-rc-alpine3.17 AS builder

COPY . /github.com/sh-latibov/telegram-bot-youtube/
WORKDIR /github.com/sh-latibov/telegram-bot-youtube/

RUN go mod download
RUN go build -o ./bin/bot cmd/bot/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=0 /github.com/sh-latibov/telegram-bot-youtube/bin/bot .
COPY --from=0 /github.com/sh-latibov/telegram-bot-youtube/configs configs/

EXPOSE 80

CMD ["./bot"]