package routes

import (
	"final-project-golang-individu/config"
	"final-project-golang-individu/controllers"
	"final-project-golang-individu/middlewares"
	"final-project-golang-individu/repositories"
	"final-project-golang-individu/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Middleware CORS
	corsConfig := cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
	r.Use(cors.New(corsConfig))

	// Repositories
	userRepository := repositories.NewUserRepository(config.DB)
	roleRepository := repositories.NewRoleRepository(config.DB)
	newsRepository := repositories.NewNewsRepository(config.DB)
	commentRepository := repositories.NewCommentRepository(config.DB)

	// Services
	userService := services.NewUserService(userRepository)
	authService := services.NewAuthService(userRepository, roleRepository)
	newsService := services.NewNewsService(newsRepository)
	commentService := services.NewCommentService(commentRepository)

	// Controllers
	authController := controllers.NewAuthController(authService)
	userController := controllers.NewUserController(userService)
	newsController := controllers.NewNewsController(newsService)
	commentController := controllers.NewCommentController(commentService)

	// Routes
	r.POST("/register", authController.Register)
	r.POST("/login", authController.Login)

	// JWT Protected routes
	auth := r.Group("/")
	auth.Use(middlewares.JWTAuthMiddleware())
	{
		auth.PUT("/change-password", userController.ChangePassword)
		auth.GET("/profile", userController.GetProfile)
		auth.PUT("/profile", userController.UpdateProfile)
		auth.GET("/users", middlewares.AuthorizeRole("admin"), userController.GetAllUsers)
		auth.GET("/users/:id", middlewares.AuthorizeRole("admin"), userController.GetUserByID)

		// News routes
		auth.POST("/news", middlewares.AuthorizeRole("admin", "editor"), newsController.CreateNews)
		auth.PUT("/news/:id", middlewares.AuthorizeRole("admin", "editor"), newsController.UpdateNews)
		auth.DELETE("/news/:id", middlewares.AuthorizeRole("admin"), newsController.DeleteNews)
		auth.GET("/news/:id", newsController.GetNews)
		auth.GET("/news", newsController.GetAllNews)

		// Comment routes
		auth.POST("/comments", commentController.CreateComment)
		auth.GET("/comments/:id", commentController.GetComment)
		auth.GET("/news/:id/comments", commentController.GetCommentsByNews) // Changed from :news_id to :id
		auth.PUT("/comments/:id", commentController.UpdateComment)
		auth.DELETE("/comments/:id", commentController.DeleteComment)
	}

	return r
}
