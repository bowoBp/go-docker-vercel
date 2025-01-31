# Gunakan base image Golang untuk build
FROM golang:1.20-alpine AS builder

# Set working directory
WORKDIR /app

# Copy go.mod dan go.sum untuk caching dependencies
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download && go mod verify

# Copy seluruh kode project ke dalam container
COPY . .

# Build aplikasi Go
RUN go build -o main .

# Tahap kedua: Production
FROM alpine:latest

# Set working directory
WORKDIR /root/

# Copy binary hasil build dari tahap sebelumnya
COPY --from=builder /app/main .

# Pastikan ENV PORT sudah ada
ENV PORT=8080
EXPOSE 8080

# Jalankan aplikasi dan tambahkan log agar terlihat di runtime logs
CMD ["sh", "-c", "echo 'Starting Docker container...' && ./main"]
