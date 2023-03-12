# Import the base image golang:alpine3.17
FROM golang:alpine3.17
# Set the Current Working Directory inside the container
WORKDIR /absolut-music
# Copy the source files from the current directory to the Working Directory inside the container
COPY . .
# Build the Go app
RUN go build -o ./bin/absolut-music
# Command being executed when starting the container
CMD [ "bin/absolut-music" ]