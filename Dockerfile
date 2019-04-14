FROM golang:latest


COPY . .

RUN go run main.go