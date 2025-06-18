FROM golang:1.23-alpine 

WORKDIR /user

COPY go.mod go.sum ./
COPY ./ ./

RUN go mod tidy

RUN go build -o user ./cmd/main.go

CMD ["./user"]