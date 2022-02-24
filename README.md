# rng

[![CI](https://github.com/acchiao/rng/actions/workflows/ci.yml/badge.svg)](https://github.com/acchiao/rng/actions/workflows/ci.yml)

![return 4](https://www.explainxkcd.com/wiki/images/f/fe/random_number.png)

## Prerequisites

- [Go] ^1.17

[go]: https://go.dev/

## Quickstart

```
go build ./...
go test ./...
go mod tidy

docker build --file Dockerfile --tag rng --load .
```
