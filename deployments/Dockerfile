FROM golang:1.17.1-alpine3.14 as builder

WORKDIR /api/cmd

COPY . /api

RUN go build -o main .

FROM alpine:latest

COPY --from=builder api/cmd/main /api/cmd/main

ENTRYPOINT ["api/cmd/main"]