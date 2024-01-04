FROM golang:1.21.5-alpine

WORKDIR /server

RUN go install github.com/cosmtrek/air@latest

WORKDIR /server

CMD ["air"]
