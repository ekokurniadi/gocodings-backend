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
	webhandler "web-portfolio-backend/web/handler"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// godotenv not used for deploy on heroku
	// env := godotenv.Load()
	// if env != nil {
	// 	log.Fatal("Error loading .env file")
	// }
	host := os.Getenv("DB_HOST")
	userHost := os.Getenv("DB_USER")
	userPass := os.Getenv("DB_PASSWORD")
	databaseName := os.Getenv("DB_DATABASE")
	databasePort := os.Getenv("DB_PORT")

	dsn := "host=" + host + " user=" + userHost + " password=" + userPass + " dbname=" + databaseName + " port=" + databasePort + " sslmode=require TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return
	}
	db.AutoMigrate(&schema.Portofolio{})
	fmt.Println("Database Connected")

	aboutRepository := repository.NewAboutRepository(db)
	aboutService := service.NewAboutService(aboutRepository)
	aboutHandler := handler.NewAboutHandler(aboutService)

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	authService := middleware.NewService()
	userHandler := handler.NewUserHandler(userService, authService)

	portfolioRepository := repository.NewPortofolioRepository(db)
	portfolioService := service.NewPortfolioService(portfolioRepository)

	sessionWebHandler := webhandler.NewSessionHandler(userService)
	userWebHandler := webhandler.NewUserHandler(userService)
	aboutWebHandler := webhandler.NewAboutHandler(aboutService)
	portfolioWebHandler := webhandler.NewPortfolioHandler(portfolioService)

	router := gin.Default()
	router.Use(cors.Default())
	cookieStore := cookie.NewStore([]byte(middleware.SECRET_KEY))
	router.Use(sessions.Sessions("bwastartup", cookieStore))
	router.LoadHTMLGlob("web/templates/**/*")
	router.Static("css", "./web/assets/css")
	router.Static("js", "./web/assets/js")
	router.Static("webfonts", "./web/assets/webfonts")
	router.Static("images", "./images")
	api := router.Group("/api/v1")

	api.GET("/abouts", aboutHandler.GetAbouts)
	api.GET("/abouts/:id", authMiddleware(authService, userService), aboutHandler.GetAbout)

	api.GET("/users", userHandler.GetUsers)
	api.PUT("/users/:id", userHandler.Update)
	api.DELETE("/users/:id", userHandler.Delete)
	api.POST("/users", userHandler.Create)
	api.POST("/login", userHandler.Login)

	router.GET("/", sessionWebHandler.New)
	router.GET("/dashboard", authAdminMiddleWare(), sessionWebHandler.Dashboard)
	router.GET("/logout", sessionWebHandler.Logout)
	router.POST("/sessions", sessionWebHandler.LoginAction)

	router.GET("/users", authAdminMiddleWare(), userWebHandler.Index)
	router.GET("/users/new", authAdminMiddleWare(), userWebHandler.New)
	router.GET("/users/update/:id", authAdminMiddleWare(), userWebHandler.Update)
	router.GET("/users/delete/:id", authAdminMiddleWare(), userWebHandler.Delete)
	router.POST("/users", authAdminMiddleWare(), userWebHandler.Create)
	router.POST("/users/update_action/:id", authAdminMiddleWare(), userWebHandler.UpdateAction)

	router.GET("/abouts", authAdminMiddleWare(), aboutWebHandler.Index)

	router.GET("/portfolios", authAdminMiddleWare(), portfolioWebHandler.Index)
	router.GET("/portfolios/new", authAdminMiddleWare(), portfolioWebHandler.New)
	router.POST("/portfolios", authAdminMiddleWare(), portfolioWebHandler.Create)
	router.Run()
}

func authAdminMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)

		userIDSession := session.Get("userID")
		if userIDSession == nil {
			c.Redirect(http.StatusFound, "/")
			return
		}
	}
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
