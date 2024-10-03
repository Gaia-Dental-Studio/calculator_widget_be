# Menggunakan image Go resmi sebagai base image
FROM golang:1.22

# Set working directory di dalam container
WORKDIR /app

# Salin go.mod dan go.sum untuk menginstall dependencies
COPY go.mod go.sum ./

# Install dependencies
RUN go mod download

# Salin seluruh kode sumber ke dalam working directory
COPY . .

# Build aplikasi Go
RUN go build -o myapp .

# Ekspos port yang digunakan aplikasi
EXPOSE 8080

# Jalankan aplikasi
CMD ["./myapp"]
