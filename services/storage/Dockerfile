FROM golang:1.24.3-alpine 

WORKDIR /storage

COPY go.mod go.sum ./
COPY ./ ./

RUN go mod tidy

RUN go build -o storage ./cmd/main.go

CMD ["./storage"]