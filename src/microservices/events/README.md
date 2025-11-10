# Proxy

## Prerequisites

- Docker

## Getting Started

### Option 1: Using Docker (Recommended)

The easiest way to start the application is to use Docker:


```bash
docker build -t events:latest .
docker run -d -p 8081:8081 events:latest 
```

### Option 2: Native run (Golang required)

If you prefer to run the application without Docker, run following command from `src` dir:

```bash
go run *.go
```
