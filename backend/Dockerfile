FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o bin/incident-service cmd/service/main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/bin/incident-service .

RUN chmod +x incident-service

EXPOSE 8080

CMD ["./incident-service"]
