# Use the official Golang image as the base image
FROM golang:1.22

# Set working directory di dalam container
WORKDIR /app

# Salin kode aplikasi ke dalam container
COPY . .

# Install dependencies (jika ada) dan build aplikasi Go
# RUN go build -o myapp .

# Tentukan perintah yang dijalankan saat container dimulai
CMD ["go", "run", "."]
