FROM golang:1.25

WORKDIR /app

COPY . .

RUN go build -o main ./cmd/main.go

EXPOSE 8000

CMD ["./main"]