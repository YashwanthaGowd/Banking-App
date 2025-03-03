FROM golang:1.24-alpine as builder

WORKDIR /app

# Copy protos and initialize its module
COPY protos/ /protos/
WORKDIR /protos

# Switch back to app directory and copy service code
WORKDIR /app
COPY account-service/ .

# Download dependencies and build
RUN go mod tidy
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Final stage
FROM alpine:latest
RUN apk add --no-cache ca-certificates
WORKDIR /app

# Copy the binary and protos
COPY --from=builder /app/main .
COPY --from=builder /protos/ ./protos/
COPY --from=builder /app/config /app/config
COPY --from=builder /app/resources /app/resources

CMD ["./main"]