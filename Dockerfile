# Birinci aşama: Uygulamayı derleme aşaması
FROM golang:1.19 AS builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -ldflags="-s -w" -o apiserver .

# İkinci aşama: Minimal çalıştırılabilir imaj oluşturma
FROM alpine:latest

WORKDIR /

COPY --from=builder /build/apiserver .
COPY --from=builder /build/.env .

ENTRYPOINT ["./apiserver"]