FROM golang:1.23-alpine 

WORKDIR /api-gateway

COPY go.mod go.sum ./
COPY ./ ./

RUN go mod tidy

RUN go build -o api-gateway ./cmd/main.go

CMD ["./api-gateway"]