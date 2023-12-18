package routes

import (
	"sanberexp/controller"
	"sanberexp/middlewares"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// set db to gin context
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	r.POST("/register", controller.Register)
	r.POST("/login", controller.Login)

	r.GET("/movies", controller.GetAllMovie)
	r.GET("/:id", controller.GetMovieById)

	moviesMiddlewareRoute := r.Group("/movies")
	moviesMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
	moviesMiddlewareRoute.POST("/", controller.CreateMovie)
	moviesMiddlewareRoute.PATCH("/:id", controller.UpdateMovie)
	moviesMiddlewareRoute.DELETE("/:id", controller.DeleteMovie)

	r.GET("/age-rating-categories", controller.GetAllRating)
	r.GET("/age-rating-categories/:id", controller.GetRatingById)
	r.GET("/age-rating-categories/:id/movies", controller.GetMoviesByRatingId)

	ratingMiddlewareRoute := r.Group("/age-rating-categories")
	ratingMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
	ratingMiddlewareRoute.POST("/", controller.CreateRating)
	ratingMiddlewareRoute.PATCH("/:id", controller.UpdateRating)
	ratingMiddlewareRoute.DELETE("/:id", controller.DeleteRating)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
