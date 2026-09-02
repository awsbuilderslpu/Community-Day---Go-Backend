package main

import (
	"log"
	"net/http"

	"github.com/awsbuilderslpu/Community-Day---Go-Backend/internal/config"
	"github.com/awsbuilderslpu/Community-Day---Go-Backend/internal/database"
	"github.com/awsbuilderslpu/Community-Day---Go-Backend/internal/module/auth"
)

func main() {
	cfg := config.LoadConfig()

	db, err := database.InitDB(cfg.DBDSN)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Auto-migrate the User schema
	err = db.AutoMigrate(&auth.User{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Initialize Auth Module
	authRepo := auth.NewRepository(db)
	authService := auth.NewService(authRepo, cfg.JWTSecret)
	authHandler := auth.NewHandler(authService)

	mux := http.NewServeMux()

	// Register Auth Routes
	authHandler.RegisterRoutes(mux)

	// Health check route
	mux.HandleFunc("/api/v1/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"success","message":"SCD Event Management System Backend is running smoothly!"}`))
	})

	log.Printf("Server starting on port %s...\n", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, mux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
