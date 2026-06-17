FROM golang:1.26.4-alpine AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/main.go

FROM alpine:latest
WORKDIR /app

# Timezone setting
RUN apk --no-cache add tzdata
ENV TZ=Asia/Seoul

COPY --from=builder /app/main .

CMD ["./main"]