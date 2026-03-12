FROM golang:1.24-alpine AS builder
WORKDIR /app

COPY shared/ ../shared/

COPY calculation_history_api/ .
RUN go mod download
RUN go build -o calculation_history_api .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/calculation_history_api .

EXPOSE 8081
CMD ["./calculation_history_api"]