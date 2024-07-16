FROM golang:alpine

RUN apk update && apk add --no-cache git

WORKDIR /home

COPY . .

RUN go mod tidy

RUN go build -o app

ENTRYPOINT ["/home/app"]
