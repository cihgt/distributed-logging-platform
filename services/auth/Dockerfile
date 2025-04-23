FROM golang:1.24

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

WORKDIR /app/services/auth
ENV GOPROXY=direct
RUN go mod tidy && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o auth

CMD ["./auth"]
