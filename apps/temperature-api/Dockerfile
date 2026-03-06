FROM golang:1.22-alpine AS builder

WORKDIR /app
COPY . .
RUN go build -o temperature-api main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/temperature-api .

EXPOSE 8081
CMD ["./temperature-api"]
