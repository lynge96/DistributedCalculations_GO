FROM golang:1.26.1-alpine AS builder
WORKDIR /app

COPY shared/ ../shared/

COPY authenticator_api/ .
RUN go mod download
RUN go build -o authenticator_api .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/authenticator_api .

EXPOSE 8082
CMD ["./authenticator_api"]