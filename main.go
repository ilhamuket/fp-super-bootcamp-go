package main

import (
	"final-project-golang-individu/config"
	"final-project-golang-individu/docs"
	"final-project-golang-individu/routes"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// @title News API
// @version 1.0
// @description This is a sample server for a news application.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Set up Swagger documentation
	docs.SwaggerInfo.Title = "News API"
	docs.SwaggerInfo.Description = "This is a sample server for a news application."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	// Initialize database
	config.InitDB()

	// Initialize router
	r := routes.SetupRouter()

	// Run server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port jika tidak ada environment variable PORT
	}
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
