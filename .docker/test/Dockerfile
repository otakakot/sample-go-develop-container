FROM golang:1.21.5-alpine

WORKDIR /test

RUN go install github.com/smartystreets/goconvey@latest

CMD ["goconvey", "-host", "0.0.0.0"]
