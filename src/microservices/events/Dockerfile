FROM golang:alpine AS builder
WORKDIR /build
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY /src .
RUN CGO_ENABLED=0 GOOS=linux go build -o events *.go

FROM alpine:latest
WORKDIR /app

COPY --from=builder /build/events .
# EXPOSE 8081
CMD ["./events"]