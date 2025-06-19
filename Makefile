build:
	docker-compose build rago

run:
	docker-compose up rago

test-api-gateway:
	cd api-gateway && go test -v -race -vet=off ./...

test-generator:
	cd services/generator && go test -v -race -vet=off ./...

test-splitter:
	cd services/splitter && go test -v -race -vet=off ./...

test-storage:
	cd services/storage && go test -v -race -vet=off ./...

test-user:
	cd services/user && go test -v -race -vet=off ./...