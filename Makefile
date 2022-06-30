BINARY_NAME=rng

all: build test

build:
	GOARCH=amd64 GOOS=darwin go build -trimpath -tags "netgo nomsgpack" -ldflags "-s -w" -o ./bin/${BINARY_NAME}-darwin main.go
	docker buildx build --file Dockerfile --tag acchiao/rng --push .

run:
	./bin/${BINARY_NAME}
.PHONY: run

test:
	go test -cover -race -v ./... -coverprofile coverage.out -timeout 30s
.PHONY: test

clean:
	go clean
	rm ./bin/${BINARY_NAME}-darwin
.PHONY: clean
