FROM golang:latest

WORKDIR /go/src/github.com/voteright

COPY . .

RUN ls

RUN pwd

RUN go run main.go