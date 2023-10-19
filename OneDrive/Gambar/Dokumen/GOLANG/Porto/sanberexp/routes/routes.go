package routes

import (
	"sanberexp/controller"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"

	swaggerFiles "github.com/swaggo/files"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// set db to gin context
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})
	r.GET("/movies", controller.GetAllMovie)
	r.POST("/movies", controller.CreateMovie)
	r.GET("/movies/:id", controller.GetMovieById)
	r.PATCH("/movies/:id", controller.UpdateMovie)
	r.DELETE("movies/:id", controller.DeleteMovie)

	r.GET("/age-rating-categories", controller.GetAllRating)
	r.POST("/age-rating-categories", controller.CreateRating)
	r.GET("/age-rating-categories/:id", controller.GetRatingById)
	r.GET("/age-rating-categories/:id/movies", controller.GetMoviesByRatingId)
	r.PATCH("/age-rating-categories/:id", controller.UpdateRating)
	r.DELETE("age-rating-categories/:id", controller.DeleteRating)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
