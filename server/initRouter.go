package server

import (
	"codeid.northwind/controllers"
	"github.com/gin-gonic/gin"
)

func InitRouter(router *gin.Engine, controllerMgr *controllers.ControllerManager) *gin.Engine {

	categoryRoute := router.Group("/category")
	{
		// Router endpoint
		categoryRoute.GET("/", controllerMgr.CategoryController.GetListCategory)
		categoryRoute.GET("/:id", controllerMgr.CategoryController.GetCategory)
		categoryRoute.POST("/", controllerMgr.CategoryController.CreateCategory)
		categoryRoute.PUT("/:id", controllerMgr.CategoryController.UpdateCategory)
		categoryRoute.DELETE("/:id", controllerMgr.CategoryController.DeleteCategory)
	}
	return router
}
