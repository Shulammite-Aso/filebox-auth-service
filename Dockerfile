# syntax=docker/dockerfile:1

FROM golang:1.18-alpine AS build

WORKDIR /app

COPY . ./

RUN go mod download

RUN go build -o /fb-auth-service

FROM scratch

WORKDIR /

COPY --from=build /fb-auth-service /fb-auth-service

EXPOSE 50051

CMD [ "/fb-auth-service" ]