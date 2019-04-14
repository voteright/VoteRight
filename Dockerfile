FROM golang:latest

COPY . .

RUN ls

RUN go run main.go