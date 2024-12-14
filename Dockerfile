# menggunakan official base image golang
FROM golang:1.20-alpine

#set working di dalam container
WORKDIR /app

# copy go.mod dan go.sum (jik ada) sebelum melakukan go mod download (untuk caching)
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# Copy seluruh kode
COPY . .

# build aplikasi
RUN go build -o main .

#expose port 8080 di container
EXPOSE 8080

#jalankan aplikasi
CMD ["./main"]