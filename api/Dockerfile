# This Dockerfile assembles an image with all the dependencies required 
# to run golang from the command-line

# Build the Golang environment
FROM golang:1.20.2

WORKDIR /usr/src/app

RUN go mod init github.com/salace-airline/animalpocketresources

RUN go install github.com/swaggo/swag/cmd/swag@latest

RUN swag init

COPY . .
RUN go mod tidy

RUN go install github.com/cosmtrek/air@latest