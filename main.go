package main

import (
	// "context"
	"log"
	// "net/http"
	"os"

	"codeid.northwind/config"
	// "codeid.northwind/repositories"
	"codeid.northwind/server"

	_ "github.com/lib/pq"
)

func main() {
	log.Println("starting northwind restapi")

	log.Println("initializing configuration")
	config := config.InitConfig(getConfigFileName())

	log.Println("initializing database")
	dbHandler := server.InitDatabase(config)
	//log.Println(dbHandler)

	log.Println("Initializing Http Server")
	httpServer := server.InitHttpServer(config, dbHandler)

	httpServer.Start()

	// test insert to category
	// 	ctx := context.Background() // langsung membuat goroutine
	// 	queries := repositories.New(dbHandler)

	// 	newCategory, err := queries.CreateCategory(ctx,
	// 		repositories.CreateCategoryParams{
	// 			CategoryID:   101,
	// 			CategoryName: "Mainan",
	// 			Description:  "Mainan Anak",
	// 			Picture:      nil,
	// 		},
	// 	)

	// 	if err != nil {
	// 		log.Fatal("Error:", err)
	// 	}

	// 	log.Println(newCategory)
}

func getConfigFileName() string {
	env := os.Getenv("ENV")

	if env != "" {
		return "nortwind" + env
	}
	return "northwind"
}
