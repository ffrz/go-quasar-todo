# Stage 1: Build
FROM golang:1.25-alpine AS builder
RUN apk add --no-cache gcc musl-dev sqlite-dev bash git
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o tmp/main ./cmd/server

# Stage 2: Runtime minimal
FROM alpine:3.18
RUN apk add --no-cache ca-certificates sqlite bash
WORKDIR /app
COPY --from=builder /app/tmp/main ./tmp/main
EXPOSE 8080
ENV PORT=8080
ENV DB_PATH=/app/data/todo.db
CMD ["./tmp/main"]
