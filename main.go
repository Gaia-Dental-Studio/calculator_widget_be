package main

import (
	"Gaia-Dental-Studio/calculator_widget_be/controller"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// CORS Middleware
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight request (OPTIONS)
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r) // Continue to the next middleware or handler
	})
}

func main() {
	r := mux.NewRouter()

	// Apply CORS middleware globally
	//  r.Use(corsMiddleware)

	// Create a subrouter with prefix "api/v0.0.1/"
	api := r.PathPrefix("/api/v0.0.1").Subrouter()

	// Define routes under the prefix
	api.HandleFunc("/create-product", controller.StoreProduct).Methods("POST")
	//api.HandleFunc("/users/{id}", getUser).Methods("GET")
	// Set up CORS
	cors := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:3000"}), // Ganti dengan domain frontend Anda
		handlers.AllowedMethods([]string{"POST", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type"}),
	)
	// Start the server
	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", cors(r)))
	//log.Fatal(http.ListenAndServe(":8080", r))
}
