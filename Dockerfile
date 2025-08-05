FROM --platform=$BUILDPLATFORM tonistiigi/xx:1.6.1 AS xx

FROM --platform=$BUILDPLATFORM golang:1.24-alpine AS builder
COPY --from=xx / /

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY main.go ./

ARG TARGETPLATFORM
RUN --mount=type=cache,target=/root/.cache \
    CGO_ENABLED=0 xx-go build -ldflags='-w -s' -trimpath -o server main.go

FROM alpine:3.22

COPY --from=builder /app/server /usr/local/bin/server

WORKDIR /app
COPY images images

EXPOSE 8080

CMD ["/usr/local/bin/server"]
