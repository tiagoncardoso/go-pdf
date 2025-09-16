FROM golang:1.24.7 AS builder

WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o webs_pdf ./cmd/webserver/main.go

FROM debian:bullseye-slim

WORKDIR /app
# Install wkhtmltopdf, fonts, and CA certificates for HTTPS
RUN apt-get update \
    && apt-get install -y --no-install-recommends wkhtmltopdf fonts-dejavu-core ca-certificates \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/webs_pdf .

EXPOSE 8081

ENTRYPOINT ["/app/webs_pdf"]