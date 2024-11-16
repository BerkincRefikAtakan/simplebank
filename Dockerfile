# Build stage
FROM golang:1.23.2-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

# Run stage
FROM alpine
WORKDIR /app
COPY --from=builder /app/main .
COPY  app.env .
COPY db/migration ./db/migration


EXPOSE 8080
CMD ["/app/main"]