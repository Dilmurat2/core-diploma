FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o core-monolith ./src/main.go && \
    upx --best --lzma core-monolith

FROM mcr.microsoft.com/playwright:v1.49.1
WORKDIR /root/

RUN apt-get update && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/core-monolith .

RUN npx playwright install chromium

EXPOSE 8080

CMD  ./core-monolith

