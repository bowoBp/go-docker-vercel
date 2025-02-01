# Gunakan base image golang
FROM golang:1.21

# Set working directory
WORKDIR /app

# Copy semua file proyek ke dalam container
COPY . .

# Download dependensi
RUN go mod tidy

# Build aplikasi
RUN go build -o main .

# Expose port 8080
EXPOSE 8080

# Jalankan aplikasi
CMD ["/app/main"]
