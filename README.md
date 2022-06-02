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
go mod tidy

docker buildx build --file Dockerfile --tag rng --load .
```

## API Reference

![return 4](https://www.explainxkcd.com/wiki/images/f/fe/random_number.png)
