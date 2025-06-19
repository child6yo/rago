FROM golang:1.24.3-alpine 

RUN apk add --no-cache git

WORKDIR /api-gateway

COPY go.mod go.sum ./
COPY ./ ./

RUN go mod tidy

RUN go build -o api-gateway ./cmd/main.go

CMD ["./api-gateway"]