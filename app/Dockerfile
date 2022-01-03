FROM golang:latest

LABEL maintainer="Ardhi <ardhi.rofi@gmail.com>"

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

ENV PORT 4000

RUN go build

CMD ["./chirpbird"]