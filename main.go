package main

import (
	"database/sql"
	"go-myClient/config"
	"go-myClient/handlers"
	"go-myClient/services"
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

func main() {
	cfg := config.LoadConfig()

	db, err := sql.Open("postgres", cfg.DBConn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rdb := redis.NewClient(&redis.Options{
		Addr: cfg.RedisAddress,
	})

	clientRepo := repository.NewClientRepository(db)
	clientService := services.NewClientService(clientRepo, rdb)
	clientHandler := handlers.NewClientHandler(clientService)

	e := echo.New()
	e.POST("/clients", clientHandler.CreateClient)
	e.PUT("/clients/:id", clientHandler.UpdateClient)
	e.DELETE("/clients/:id", clientHandler.DeleteClient)
	e.GET("/clients/:id", clientHandler.GetClientByID)

	e.Logger.Fatal(e.Start(":8080"))
}
