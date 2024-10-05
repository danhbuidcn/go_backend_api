FROM golang:alpine AS builder

WORKDIR /build

COPY . .

RUN go mod download

RUN go build -o backend_api ./cmd/server

FROM scratch

COPY ./config /config

COPY --from=builder /build/backend_api /

ENTRYPOINT ["/backend_api", "config/local.yaml"]
