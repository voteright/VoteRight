FROM golang:latest

WORKDIR /go/src/github.com/voteright/voteright

COPY . .

RUN go get

RUN go build .

CMD ["voteright"]