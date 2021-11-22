package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"web-portfolio-backend/handler"
	"web-portfolio-backend/helper"
	"web-portfolio-backend/middleware"
	"web-portfolio-backend/repository"
	"web-portfolio-backend/schema"
	"web-portfolio-backend/service"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// env := godotenv.Load()
	// if env != nil {
	// 	log.Fatal("Error loading .env file")
	// }
	// godotenv not used for deploy on heroku
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
	db.AutoMigrate(&schema.About{}, &schema.User{})
	fmt.Println("Database Connected")

	aboutRepository := repository.NewAboutRepository(db)
	aboutService := service.NewAboutService(aboutRepository)
	aboutHandler := handler.NewAboutHandler(aboutService)

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	authService := middleware.NewService()
	userHandler := handler.NewUserHandler(userService, authService)

	router := gin.Default()
	router.Use(cors.Default())
	router.LoadHTMLGlob("web/templates/**/*")
	router.Static("css", "./web/assets/css")
	router.Static("images", "./images")
	api := router.Group("/api/v1")

	api.GET("/abouts/:id", authMiddleware(authService, userService), aboutHandler.GetAbout)
	api.POST("/users", userHandler.Create)
	api.POST("/login", userHandler.Login)

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "auth.html", nil)
	})
	router.Run()
}

func authMiddleware(authService middleware.Service, userService service.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helper.ApiResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")

		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)

		if err != nil {
			response := helper.ApiResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			response := helper.ApiResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userID := int(claim["user_id"].(float64))

		user, err := userService.UserServiceGetByID(userID)

		if err != nil {
			response := helper.ApiResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("currentUser", user)

	}

}
