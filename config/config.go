package config

import (
	"final-project-golang-individu/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // import PostgreSQL driver
	"github.com/joho/godotenv"
	"log"
	"os"
)

var DB *gorm.DB

func InitDB() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Read environment variables
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	// Connect to the database
	var err error
	DB, err = gorm.Open("postgres", "host="+dbHost+" port="+dbPort+" user="+dbUsername+" dbname="+dbName+" password="+dbPassword+" sslmode=disable")
	if err != nil {
		log.Fatal("Failed to connect database", err)
	}
	DB.LogMode(true)
	migrate()
}

func migrate() {
	DB.AutoMigrate(&models.User{}, &models.Role{}, &models.Profile{}, &models.News{}, &models.Comment{}, &models.UserRole{})
	DB.Model(&models.Profile{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	DB.Model(&models.News{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	DB.Model(&models.Comment{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	DB.Model(&models.Comment{}).AddForeignKey("news_id", "news(id)", "CASCADE", "CASCADE")
	createDefaultRoles()
}

func createDefaultRoles() {
	var adminRole, editorRole models.Role

	// Check if 'admin' role exists
	DB.Where(models.Role{Name: "admin"}).FirstOrCreate(&adminRole)
	if adminRole.ID == 0 {
		log.Println("Created default role: admin")
	}

	// Check if 'editor' role exists
	DB.Where(models.Role{Name: "editor"}).FirstOrCreate(&editorRole)
	if editorRole.ID == 0 {
		log.Println("Created default role: editor")
	}
}
