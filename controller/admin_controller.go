package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

// Struct untuk menyimpan data produk
type Product struct {
	NameProduct string `json:"name_product"`
	Description string `json:"description"`
	Category    string `json:"category"`
	// Image //what a type for image
}

// Function untuk menangani request dan mengembalikan JSON response
func StoreProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse the multipart form
	err := r.ParseMultipartForm(10 << 20) // Limit to 10 MB
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	// Get the file
    file, header, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Error retrieving file", http.StatusBadRequest)
		return
	}
	defer file.Close()

    // Create uploads folder if it doesn't exist
    os.MkdirAll("uploads", os.ModePerm) 

    // Generate a unique filename based on the original filename and timestamp
    uniqueName := fmt.Sprintf("%s_%s", time.Now().Format("20060102150405"), header.Filename)
    out, err := os.Create(filepath.Join("uploads", uniqueName))

	if err != nil {
		http.Error(w, "Error creating file", http.StatusInternalServerError)
		return
	}
	defer out.Close()

	// Copy the uploaded file to the created file
	_, err = io.Copy(out, file)
	if err != nil {
		http.Error(w, "Error saving file", http.StatusInternalServerError)
		return
	}

	// Deklarasikan variabel untuk menyimpan produk
	var product Product

	// Log data yang diterima
	log.Printf("Received product: Name Product: %s,Description: %s,Category: %s\n", product.NameProduct, product.Description, product.Category)

	// Mengatur header response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Kembalikan response JSON
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  200,
		"message": "Product stored successfully",
		"product": product,
	})
}
