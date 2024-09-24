FROM golang:1.23-alpine AS builder

RUN apk add build-base

WORKDIR /root

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY ./cmd ./cmd
COPY ./pkg ./pkg
COPY ./migrations ./migrations
RUN CGO_ENABLED=0 GOOS=linux go build -o server cmd/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates tzdata bash
WORKDIR /root
COPY --from=builder /root/server ./server
COPY --from=builder /root/migrations ./migrations

CMD ["./server"]
