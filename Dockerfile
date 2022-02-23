# syntax=docker/dockerfile:1
FROM golang:1.17.7-bullseye AS build

ENV RNG_HOME=/opt/rng

WORKDIR ${RNG_HOME}

COPY go.mod go.sum ${RNG_HOME}/
COPY go.sum ${RNG_HOME}/
RUN go mod download

COPY main.go ${RNG_HOME}/
RUN go build -tags netgo -ldflags '-s -w' -o rng

FROM gcr.io/distroless/base-debian11

COPY --from=build /opt/rng/rng /rng

EXPOSE 3000
ENTRYPOINT ["/rng"]
