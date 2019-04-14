FROM golang:latest

COPY . .

RUN ls

RUN pwd

RUN go run main.go