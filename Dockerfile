FROM golang:1.16.3-alpine

WORKDIR /app

COPY . /app

RUN go mod download
