FROM golang:alpine3.17

ADD . /go/src/myapp

WORKDIR /go/src/myapp
RUN go build -o bin/myapp
CMD [ "bin/myapp" ]