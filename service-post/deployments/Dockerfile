FROM golang:alpine

RUN apk add --no-cache git gcc musl-dev

WORKDIR /usr/app

COPY . .

RUN go mod download && go build -o main.o ./cmd/thrift

EXPOSE 9090

ENTRYPOINT ["/usr/app/main.o"]
