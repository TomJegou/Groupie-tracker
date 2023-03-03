FROM golang:1.16-alpine

ADD . /go/src/myapp

WORKDIR /go/src/myapp
RUN go build -o bin/myapp
CMD [ "bin/myapp" ]