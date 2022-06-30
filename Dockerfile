# syntax=docker/dockerfile:1
FROM golang:1.18.3-bullseye AS build

ENV RNG_HOME=/opt/rng \
      CGO_ENABLED=0 \
      GIN_MODE=release \
      GOARCH=amd64 \
      GOOS=linux

WORKDIR ${RNG_HOME}

COPY go.mod go.sum ${RNG_HOME}/
RUN go mod download

COPY main.go ${RNG_HOME}/
RUN go build -trimpath -tags "netgo nomsgpack" -ldflags "-s -w" -o rng

FROM gcr.io/distroless/base-debian11 AS rng

COPY --from=build /opt/rng/rng /rng

EXPOSE 3000
ENTRYPOINT ["/rng"]
