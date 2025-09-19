FROM golang:1.24.7 AS builder

WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o webs_pdf ./cmd/webserver/main.go

FROM surnet/alpine-wkhtmltopdf:3.22.0-024b2b2-small

WORKDIR /app

# Install dependencies
RUN apk add --no-cache \
    ca-certificates \
    fontconfig

COPY --from=builder /app/webs_pdf .

EXPOSE 8081

ENTRYPOINT ["/app/webs_pdf"]