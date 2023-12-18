package server

import (
	"database/sql"
	"log"

	"codeid.northwind/controllers"
	"codeid.northwind/repositories"
	"codeid.northwind/services"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type HttpServer struct {
	config            *viper.Viper
	router            *gin.Engine
	ControllerManager controllers.ControllerManager
	// categoryController *controllers.CategoryController
}

func InitHttpServer(config *viper.Viper, dbHandler *sql.DB) HttpServer {

	repositoryManager := repositories.NewRepositoryManager(dbHandler)

	serviceManager := services.NewServiceManager(repositoryManager)

	controllerManager := controllers.NewControllerManager(serviceManager)

	router := gin.Default()

	InitRouter(router, controllerManager)

	return HttpServer{
		config:            config,
		router:            router,
		ControllerManager: *controllerManager,
	}
}

// Running gin HttpServer
func (hs HttpServer) Start() {
	err := hs.router.Run(hs.config.GetString("http.server_address"))

	if err != nil {
		log.Fatalf("Error while starting HTTP Server: %v", err)
	}
}
