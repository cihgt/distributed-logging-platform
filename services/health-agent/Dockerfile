FROM golang:1.24

WORKDIR /app
COPY . .

RUN go build -o agent main.go

CMD ["./agent"]
