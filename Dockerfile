ARG GO_VERSION=1
FROM golang:${GO_VERSION}-bookworm as builder

WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN go build -ldflags="-s -w" -o /gotth-top ./cmd/gotth-top/main.go


FROM debian:bookworm

COPY --from=builder /gotth-top /usr/local/bin/
CMD ["gotth-top"]
