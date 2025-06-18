FROM golang:1.23-alpine 

WORKDIR /splitter

COPY go.mod go.sum ./
COPY ./ ./

RUN go mod tidy

RUN go build -o splitter ./cmd/main.go

CMD ["./splitter"]