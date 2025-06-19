FROM golang:1.24.3-alpine 

RUN apk add --no-cache git

WORKDIR /user

COPY go.mod go.sum ./
COPY ./ ./

RUN go mod tidy

RUN go build -o user ./cmd/main.go

CMD ["./user"]