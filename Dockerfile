FROM golang:1.24.7 AS builder

WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o webs_pdf ./cmd/webserver/main.go

FROM gcr.io/distroless/static

WORKDIR /app
COPY --from=builder /app/webs_pdf .

EXPOSE 8081

ENTRYPOINT ["/app/webs_pdf"]