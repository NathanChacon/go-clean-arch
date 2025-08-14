

FROM golang:1.24.2-alpine

RUN apk add --no-cache ca-certificates && update-ca-certificates

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download


COPY . .


RUN go build -o server main.go

EXPOSE 8080


CMD ["./server"]