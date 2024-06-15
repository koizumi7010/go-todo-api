FROM golang:1.21.6

ENV TZ="Asia/Tokyo"
WORKDIR /go/src/app

COPY go.mod go.sum ./
RUN go mod download
