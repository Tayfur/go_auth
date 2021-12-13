### Stage 1
FROM golang:1.17.2-alpine AS builder

WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-s -w" -o /app/auth

### Stage 2
FROM scratch
COPY --from=builder /app/auth /bin/auth
COPY .env /app
ENTRYPOINT ["/bin/auth"]