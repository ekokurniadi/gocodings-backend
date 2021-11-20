package main

import (
	"log"
	"os"

	"web-portfolio-backend/input"
	"web-portfolio-backend/repository"
	"web-portfolio-backend/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	env := godotenv.Load()
	if env != nil {
		log.Fatal("Error loading .env file")
	}
	host := os.Getenv("DB_HOST")
	userHost := os.Getenv("DB_USER")
	userPass := os.Getenv("DB_PASSWORD")
	databaseName := os.Getenv("DB_DATABASE")
	databasePort := os.Getenv("DB_PORT")

	dsn := "host=" + host + " user=" + userHost + " password=" + userPass + " dbname=" + databaseName + " port=" + databasePort + " sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return
	}

	aboutRepository := repository.NewAboutRepository(db)
	aboutService := service.NewAboutService(aboutRepository)

	var input input.InputID

	input.ID = 3

	aboutService.AboutServiceDelete(input)

	router := gin.Default()
	router.Use(cors.Default())
	api := router.Group("/api/v1")

	api.GET("/abouts")

	router.Run()
}
