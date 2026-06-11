FROM golang:1.26.4-alpine AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

FROM alpine:latest
WORKDIR /app

# ★ 핵심: 1단계 방에 덩그러니 남겨진 리눅스용 'main' 실행 파일만 이 방으로 쏙 복사해옵니다!
COPY --from=builder /app/main .

CMD ["./main"]