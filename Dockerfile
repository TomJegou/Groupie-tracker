FROM golang:1.19.6-bullseye

WORKDIR /app

COPY go.mod ./

COPY * ./

RUN go build -o /myapp main.go

EXPOSE 80

CMD [ "/myapp" ]