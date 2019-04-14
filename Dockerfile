FROM golang:latest

COPY . .

RUN ls

RUN workdir

RUN go run main.go