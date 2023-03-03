FROM golang:1.19.6-bullseye

ADD . /go/src/myapp

WORKDIR /go/src/myapp
RUN go build -o bin/myapp
CMD [ "bin/myapp" ]