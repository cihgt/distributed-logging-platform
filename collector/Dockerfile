FROM golang:1.24 AS builder

WORKDIR /app

COPY collector/go.mod collector/go.sum ./
COPY proto ./proto
RUN go mod tidy

COPY . .

WORKDIR /app/collector

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o collector

FROM scratch

WORKDIR /root/
COPY --from=builder /app/collector/collector .

EXPOSE 50051 9090
ENTRYPOINT ["/root/collector"]
