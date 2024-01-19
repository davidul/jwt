
build:
	docker build -t jwt .

run:
	docker run --rm jwt

lint:
	docker run --rm -v $(shell pwd):/app -w /app golangci/golangci-lint:v1.55.2 golangci-lint run -v