package main

import (
	"log"
	"net/http"

	"category-api/internal/config"
	"category-api/internal/handler"
	"category-api/internal/repository"
	"category-api/internal/service"

	"github.com/gorilla/mux"
)

func main() {
	// Initialize database
	if err := config.InitDB(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer config.CloseDB()
	log.Println("Database connected successfully")

	// Initialize repositories
	categoryRepo := repository.NewCategoryRepository()
	productRepo := repository.NewProductRepository()

	// Initialize services
	categoryService := service.NewCategoryService(categoryRepo)
	productService := service.NewProductService(productRepo)

	// Initialize handlers
	categoryHandler := handler.NewCategoryHandler(categoryService)
	productHandler := handler.NewProductHandler(productService)

	// Setup router
	router := mux.NewRouter()

	// Register routes
	categoryHandler.RegisterRoutes(router)
	productHandler.RegisterRoutes(router)

	// CORS middleware
	router.Use(corsMiddleware)

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}