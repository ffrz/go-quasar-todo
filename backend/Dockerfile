# builder
FROM golang:1.21 as builder
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o /app/server ./cmd/server


FROM alpine:3.18
RUN addgroup -S app && adduser -S app -G app
COPY --from=builder /app/server /server
RUN mkdir /data && chown app:app /data
USER app
WORKDIR /
EXPOSE 8080
ENTRYPOINT ["/server"]