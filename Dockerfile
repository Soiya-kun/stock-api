FROM golang:1.20-alpine AS dev

ENV GOOS=linux
ENV GOARCH=amd64
ENV CGO_ENABLED=0
ENV GO111MODULE=on

WORKDIR /app
COPY ./go.mod ./go.sum ./
RUN apk update && \
    apk add --update --no-cache git && \
    apk add --no-cache gcc && \
    apk add --no-cache musl-dev && \
    go mod download
EXPOSE 80
