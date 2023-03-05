FROM golang:alpine3.17

WORKDIR /absolut-music
COPY . .
RUN go build -o ./bin/absolut-music
CMD [ "bin/absolut-music" ]