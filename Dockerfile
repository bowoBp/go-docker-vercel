# Gunakan base image Golang
FROM golang:1.20-alpine AS builder

# Set working directory di dalam container
WORKDIR /app

# Copy file go.mod dan go.sum untuk caching dependencies
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

# Copy binary yang telah dibangun dari stage sebelumnya
COPY --from=builder /app/main .

# Expose port agar sesuai dengan Vercel
ENV PORT=8080
EXPOSE 8080

# Jalankan aplikasi
CMD ["./main"]
