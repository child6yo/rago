FROM golang:1.24.3-alpine

RUN apk add --no-cache git

WORKDIR /splitter

COPY go.mod go.sum ./
COPY ./ ./

RUN go mod tidy

RUN go build -o splitter ./cmd/main.go

CMD ["./splitter"]