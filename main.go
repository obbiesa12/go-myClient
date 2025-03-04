package main

import (
	"database/sql"
	"go-myClient/config"
	"go-myClient/handlers"
	"go-myClient/repository"
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

	clientRepo := repository.NewClient(db)
	clientSvc := services.NewClientService(clientRepo, rdb)
	clientHdlr := handlers.NewClientHandler(cfg, clientSvc)

	e := echo.New()
	e.POST("/clients", clientHdlr.Create)
	e.PUT("/clients/:id", clientHdlr.Update)
	e.DELETE("/clients/:id", clientHdlr.Delete)
	e.GET("/clients/:id", clientHdlr.GetByID)

	e.Logger.Fatal(e.Start(":8080"))
}
