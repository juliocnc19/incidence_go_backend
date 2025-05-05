FROM golang:1.23.4-alpine AS builder

RUN apk add --no-cache gcc musl-dev

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -ldflags='-w -s' -o main .

FROM alpine:latest

WORKDIR /app

RUN addgroup -S appgroup && adduser -S appuser -G appgroup

COPY --from=builder /app/main .

RUN mkdir uploads && chown appuser:appgroup uploads

RUN chown -R appuser:appgroup /app

USER appuser

EXPOSE 3001

CMD ["./main"]