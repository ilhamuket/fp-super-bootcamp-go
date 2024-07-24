package routes

import (
	"final-project-golang-individu/config"
	"final-project-golang-individu/controllers"
	"final-project-golang-individu/middlewares"
	"final-project-golang-individu/repositories"
	"final-project-golang-individu/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"time"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"}
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	corsConfig.AllowHeaders = []string{"Content-Type", "Authorization"}
	corsConfig.ExposeHeaders = []string{"Content-Length"}
	corsConfig.AllowCredentials = true
	corsConfig.MaxAge = 12 * time.Hour

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
	userController := controllers.NewUserController(userService, roleRepository)
	newsController := controllers.NewNewsController(newsService)
	commentController := controllers.NewCommentController(commentService)

	// Serve static files
	r.Static("/public", "./public")

	// Root route to serve HTML
	r.GET("/", func(c *gin.Context) {
		c.File("./public/index.html")
	})

	// Routes
	r.POST("/register", authController.Register)
	r.POST("/login", authController.Login)
	r.GET("/news/:id", newsController.GetNews)
	r.GET("/news", newsController.GetAllNews)
	r.GET("/comments/:id", commentController.GetComment)
	r.GET("/news/comments/:news_id", commentController.GetCommentsByNews)

	// JWT Protected routes
	auth := r.Group("/")
	auth.Use(middlewares.JWTAuthMiddleware())
	{
		auth.PUT("/change-password", userController.ChangePassword)
		auth.GET("/profile", userController.GetProfile)
		auth.PUT("/profile", userController.UpdateProfile)
		auth.GET("/users", middlewares.AuthorizeRole("admin"), userController.GetAllUsers)
		auth.GET("/users/:id", middlewares.AuthorizeRole("admin"), userController.GetUserByID)
		auth.POST("/users", middlewares.AuthorizeRole("admin"), userController.CreateUser)
		auth.PUT("/users/:id", middlewares.AuthorizeRole("admin"), userController.UpdateUser)
		auth.DELETE("/users/:id", middlewares.AuthorizeRole("admin"), userController.DeleteUser)

		// News routes
		auth.POST("/news", middlewares.AuthorizeRole("admin", "editor"), newsController.CreateNews)
		auth.PUT("/news/:id", middlewares.AuthorizeRole("admin", "editor"), newsController.UpdateNews)
		auth.DELETE("/news/:id", middlewares.AuthorizeRole("admin"), newsController.DeleteNews)

		// Comment routes
		auth.POST("/comments", commentController.CreateComment)
		auth.PUT("/comments/:id", commentController.UpdateComment)
		auth.DELETE("/comments/:id", commentController.DeleteComment)

	}

	// Swagger endpoint
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
