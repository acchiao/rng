# rng

[![CI](https://github.com/acchiao/rng/actions/workflows/ci.yml/badge.svg)](https://github.com/acchiao/rng/actions/workflows/ci.yml)
[![CodeQL](https://github.com/acchiao/rng/actions/workflows/codeql.yml/badge.svg)](https://github.com/acchiao/rng/actions/workflows/codeql.yml)

## Prerequisites

- [Go] ^1.17

[go]: https://go.dev/

## Quickstart

```shell
go build ./...
go test ./...
go mod tidy

docker build --file Dockerfile --tag rng --load .
```

## API Reference

![return 4](https://www.explainxkcd.com/wiki/images/f/fe/random_number.png)
