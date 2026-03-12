FROM golang:1.26.1-alpine AS builder
WORKDIR /app

COPY shared/ ../shared/

COPY calculator_api/ .
RUN go mod download
RUN go build -o calculator_api .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/calculator_api .

EXPOSE 8080
CMD ["./calculator_api"]