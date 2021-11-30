FROM golang:1.17 as builder

ARG APP_NAME=babycrab
ARG CGO_ENABLED=0
ARG GOOS=linux
ARG GOARCH=amd64

RUN apt-get update && apt-get install -y ca-certificates && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /$APP_NAME
COPY . /$APP_NAME

RUN make build

FROM alpine:3.11

LABEL description="babycrab is a simple API"

RUN apk --no-cache add tzdata ca-certificates

COPY --from=builder /babycrab/serverd /

CMD ./serverd

# migrations
FROM migrate/migrate:v4.14.1

COPY ./db/migrations /migrations