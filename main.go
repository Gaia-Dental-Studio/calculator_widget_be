package main

import (
	"Gaia-Dental-Studio/calculator_widget_be/controller"
	"Gaia-Dental-Studio/calculator_widget_be/model"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

        w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func ConnectDatabase() {
    dsn := "host=localhost user=rivaldosetyo password=1212 dbname=medigear_db port=5432 sslmode=disable TimeZone=Asia/Jakarta"
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database!", err)
    }
    fmt.Println("Database connected successfully")

    DB.AutoMigrate(&model.Product{})
}

func main() {
	r := mux.NewRouter()
    ConnectDatabase()
    controller.DB = DB
    r.PathPrefix("/uploads/").Handler(http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads/"))))

	api := r.PathPrefix("/api/v0.0.1").Subrouter()

	api.HandleFunc("/create-product", controller.StoreProduct).Methods("POST")
	api.HandleFunc("/get-products", controller.GetProducts).Methods("GET")
	api.HandleFunc("/get-product", controller.GetProductsById).Methods("GET")

	cors := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:3000"}), // Ganti dengan domain frontend Anda
		handlers.AllowedMethods([]string{"POST", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type"}),
	)

    log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", cors(r)))
}
