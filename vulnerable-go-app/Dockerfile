FROM golang:1.21-alpine AS builder

WORKDIR /app

# Copy the module files first
COPY webapp/go.mod ./
# If you had dependencies, you'd run 'go mod download' here

# Copy the source code
COPY webapp/main.go ./

# Build the binary
RUN go build -o main main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .
RUN apk add --no-cache curl
EXPOSE 8080
CMD ["./main"]
