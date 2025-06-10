FROM golang:1.23-alpine 

WORKDIR /generator

COPY go.mod go.sum ./
COPY ./ ./

RUN go mod tidy

RUN go build -o generator ./main.go

CMD ["./generator"]

