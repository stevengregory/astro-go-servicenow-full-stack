FROM golang:1.22 as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o myapp cmd/service/main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/myapp .

RUN chmod +x myapp

EXPOSE 8080

CMD ["./myapp"]
