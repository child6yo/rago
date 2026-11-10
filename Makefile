build:
	docker compose build

up:
	docker compose up -d --build

down:
	docker compose down

logs:
	docker compose logs -f

test:
	go test -race ./...

vet:
	go vet ./...

tidy:
	go mod tidy
