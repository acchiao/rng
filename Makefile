BINARY_NAME=rng

all: clean build lint test push

build:
	GOARCH=amd64 GOOS=darwin go build -race -trimpath -tags "netgo nomsgpack" -ldflags "-s -w" -o ./bin/${BINARY_NAME}-darwin main.go
	docker buildx build --file Dockerfile --tag acchiao/rng:dev --load .
.PHONY: build

push:
	docker buildx build --file Dockerfile --tag acchiao/rng:dev --push .
.PHONY: push

run:
	./bin/${BINARY_NAME}
.PHONY: run

lint:
	gofumpt -l -w .
	golangci-lint run
.PHONY: lint

test:
	go test -tags=unit -cover -race -v ./... -coverprofile coverage.out -timeout 30s
.PHONY: test

clean:
	go clean
	rm ./bin/${BINARY_NAME}-darwin
.PHONY: clean
