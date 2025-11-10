FROM golang:alpine AS builder
WORKDIR /build
COPY /src .
RUN CGO_ENABLED=0 GOOS=linux go build -o proxy *.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /build/proxy .
# EXPOSE 8081
CMD ["./proxy"]