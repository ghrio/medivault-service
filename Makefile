serve:
	go run cmd/api/main.go
test:
	go test -v ./...
test-coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out
build:
	go build -o bin/medivault-service cmd/api/main.go
run:
	./bin/medivault-service
docker-build:
	docker build -t medivault-service .
docker-run:
	docker run -p 8080:8080 medivault-service
docker-compose-up:
	docker-compose up
docker-compose-down:
	docker-compose down
swagger:
	swag init -g cmd/server/server.go --parseDependency --parseInternal
test-container:
.PHONY: serve test test-coverage build run docker-build docker-run docker-compose-up docker-compose-down
