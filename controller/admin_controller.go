package controller

import (
	"Gaia-Dental-Studio/calculator_widget_be/helper"
	"Gaia-Dental-Studio/calculator_widget_be/model"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

var DB *gorm.DB

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

    nameProduct := r.FormValue("name_product")
    description := r.FormValue("description")
    category := r.FormValue("category")
    price := r.FormValue("price")
    freeWarranty := r.FormValue("free_warranty")

	// Get the file
	file, header, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Error retrieving image", http.StatusBadRequest)
		return
	}
	defer file.Close()

    imagePath, err := helper.UploadFile(file, header, "uploads/images")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Get the file
	filePdf, headerPdf, errPdf := r.FormFile("pdf")
	if errPdf != nil {
		http.Error(w, "Error retrieving pdf", http.StatusBadRequest)
		return
	}
	defer filePdf.Close()

    pdfPath, err := helper.UploadFile(filePdf, headerPdf, "uploads/pdf")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

	product := &model.Product{
		NameProduct: nameProduct,
		Description: description, 
        Category: category,
        Image: imagePath,
        Document: pdfPath,
        Price: price,
        FreeWarranty: freeWarranty,
    }

	result := DB.Create(&product)
	if result.Error != nil {
		http.Error(w, "Error saving data", http.StatusInternalServerError)
		return
	}
	log.Printf("Received product: Name Product: %s,Description: %s,Category: %s\n", product.NameProduct, product.Description, product.Category)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  200,
		"message": "Product stored successfully",
		"product": product,
	})
}

const BaseURL = "http://localhost:8080/"

func GetProducts(w http.ResponseWriter, r *http.Request) {
	var products []model.Product

	result := DB.Find(&products)
	if result.Error != nil {
		http.Error(w, "Failed to fetch products", http.StatusInternalServerError)
		return
	}

	for i, product := range products {
		products[i].Image = BaseURL + product.Image
		products[i].Document = BaseURL + product.Document
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func GetProductsById(w http.ResponseWriter, r *http.Request) {
    // Get the ID from query parameters
	idStr := r.URL.Query().Get("id")

	// Convert ID from string to integer
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	// Select product by ID from the database
	var product model.Product
	result := DB.First(&product, id)
	if result.Error != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	// Return product as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}
