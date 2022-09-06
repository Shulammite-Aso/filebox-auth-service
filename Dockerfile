# syntax=docker/dockerfile:1

FROM golang:1.18-alpine AS build

WORKDIR /app

COPY . ./

RUN go mod download

EXPOSE 50051

ENTRYPOINT go run cmd/main.go