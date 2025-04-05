# 1. Build Stage
FROM golang:1.21 AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o main .

# 2. Run Stage
FROM gcr.io/distroless/base-debian11

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]
