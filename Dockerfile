FROM golang:1.8-alpine

RUN apk --update add git #&& \
    # go get \
    #   github.com/golang/dep

RUN mkdir -p /go/src/github.com/minodisk/go-learn-gorm
WORKDIR /go/src/github.com/minodisk/go-learn-gorm
COPY . .

CMD go run main.go
