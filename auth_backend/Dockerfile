FROM golang:1.24.2-alpine AS builder

LABEL maintainer="Devesh Choudhary" \
    version="1.0.0" \
    description="User Authentication Service"

RUN apk update && \
    apk add --no-cache gcc musl-dev && \
    rm -rf /var/cache/apk/*

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build \
    -ldflags="-w -s -X main.version=1.0.0" \
    -trimpath \
    -o userAuth ./cmd/main.go

FROM alpine:latest

LABEL maintainer="Devesh Choudhary" \
    version="1.0.0" \
    description="User Authentication Service"

RUN apk update && \
    apk add --no-cache ca-certificates tzdata curl wget && \
    update-ca-certificates && \
    rm -rf /var/cache/apk/*

WORKDIR /app

COPY --from=builder /app/userAuth .

ENV GIN_MODE=release \
    TZ=UTC

EXPOSE 8080

ENTRYPOINT ["./userAuth"]
