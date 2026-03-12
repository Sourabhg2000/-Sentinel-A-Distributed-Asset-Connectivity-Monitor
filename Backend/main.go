package main

import (
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"Golang_Test/config"
	"Golang_Test/routes"
)

func main() {
	// 1. Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// 2. Initialize MongoDB connection
	config.ConnectDB()
	defer config.DisconnectDB()

	// 3. Initialize Gin router
	router := gin.Default()

	// 4. CORS Configuration
	// Note: renamed variable to 'corsConfig' to avoid conflict with 'config' package
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:5173"}
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}

	// CRITICAL: Added "Authorization" and "Accept" to allow JWT tokens
	corsConfig.AllowHeaders = []string{
		"Origin",
		"Content-Length",
		"Content-Type",
		"Authorization",
		"Accept",
	}

	// Allow browser to send credentials (cookies, auth headers)
	corsConfig.AllowCredentials = true

	// Apply CORS middleware before routes
	router.Use(cors.New(corsConfig))

	// 5. Setup routes
	routes.SetupRoutes(router)

	// 6. Get port from environment or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	// 7. Start server
	log.Printf("Server starting on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
