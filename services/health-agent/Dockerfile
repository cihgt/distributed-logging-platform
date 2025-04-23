FROM golang:1.24

WORKDIR /app
COPY services/health-agent/go.mod ./services/health-agent/go.mod
COPY services/health-agent/go.sum ./services/health-agent/go.sum
COPY services/health-agent/ ./services/health-agent/
COPY ../../proto ./proto

WORKDIR /app/services/health-agent

RUN go mod tidy

RUN go build -o agent main.go

CMD ["./agent"]
