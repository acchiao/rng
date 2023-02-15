# syntax=docker/dockerfile:1
FROM golang:1.20.1-bullseye AS build

ENV RNG_HOME=/opt/rng \
      CGO_ENABLED=0 \
      GIN_MODE=release \
      GOARCH=amd64 \
      GOOS=linux

WORKDIR ${RNG_HOME}

COPY go.mod go.sum ${RNG_HOME}/
RUN go mod download

COPY main.go ${RNG_HOME}/
RUN go build -trimpath -tags "netgo nomsgpack" -ldflags "-s -w" -o ${RNG_HOME}/bin/rng main.go

FROM gcr.io/distroless/base-debian11 AS rng

ENV RNG_HOME=/opt/rng \
      CGO_ENABLED=0 \
      GIN_MODE=release \
      GOARCH=amd64 \
      GOOS=linux

COPY --from=build ${RNG_HOME}/bin/rng /rng

EXPOSE 3000
ENTRYPOINT ["/rng"]
