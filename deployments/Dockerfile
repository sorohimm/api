FROM golang:latest as builder

WORKDIR /User/sorohimm/GoProjects/api/cmd

COPY . /User/sorohimm/GoProjects/api

RUN go build -o main .

FROM ubuntu:latest

COPY --from=builder /User/sorohimm/GoProjects/api/cmd/main /User/sorohimm/GoProjects/api/cmd/main

ENTRYPOINT ["/User/sorohimm/GoProjects/api/cmd"]


