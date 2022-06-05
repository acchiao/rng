# rng

[![CI](https://github.com/acchiao/rng/actions/workflows/ci.yml/badge.svg)](https://github.com/acchiao/rng/actions/workflows/ci.yml)
[![Release](https://github.com/acchiao/rng/actions/workflows/release.yml/badge.svg)](https://github.com/acchiao/rng/actions/workflows/release.yml)

## Prerequisites

- [Go] ^1.18

[go]: https://go.dev/

## Quickstart

```shell
go get ./...
go build ./...
go test ./...
go fmt ./...
go mod tidy

docker buildx build --file Dockerfile --tag rng --load .
docker run --name rng --detach --publish 3000:3000 rng
curl http://127.0.0.1:3000
docker stop rng
docker rm rng
```

## API Reference

![return 4](https://www.explainxkcd.com/wiki/images/f/fe/random_number.png)
