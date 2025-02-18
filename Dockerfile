ARG GO_VERSION=1
FROM golang:${GO_VERSION}-bookworm AS builder

WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN go build -v -o /run-app ./cmd/web


FROM debian:bookworm

RUN apt-get update && apt-get install -y ca-certificates

COPY --from=builder /run-app /usr/local/bin/
COPY /ui/static /ui/static
CMD ["run-app"]
