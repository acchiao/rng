# syntax=docker/dockerfile:1
FROM golang:1.17.8-bullseye AS build

ENV LIMELIGHT_HOME=/opt/limelight \
      CGO_ENABLED=0 \
      GIN_MODE=release \
      GOARCH=amd64 \
      GOOS=linux

WORKDIR ${LIMELIGHT_HOME}

COPY go.mod go.sum ${LIMELIGHT_HOME}/
RUN go mod download

COPY limelight.go ${LIMELIGHT_HOME}/
RUN go build -tags "netgo nomsgpack" -ldflags "-s -w" -o limelight

FROM gcr.io/distroless/base-debian11 AS limelight

COPY --from=build /opt/limelight/limelight /limelight

EXPOSE 3000
ENTRYPOINT ["/limelight"]
