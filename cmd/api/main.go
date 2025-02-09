package main

import (
	"go_role/internal/config"
	"go_role/internal/core/services"
	"go_role/internal/handlers"
	"go_role/internal/infrastructure/database"
	"go_role/internal/repositories/postgres"

	"github.com/labstack/echo/v4"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	db, err := database.NewPostgresDB(cfg.Database)
	if err != nil {
		panic(err)
	}

	postRepo := postgres.NewPostRepository(db)
	userRepo := postgres.NewUserRepository(db)

	postService := services.NewPostService(postRepo, userRepo)
	postHandler := handlers.NewPostHandler(postService)

	e := echo.New()

	api := e.Group("/api")
	{
		posts := api.Group("/post")
		{
			posts.POST("", postHandler.Create)
			posts.GET("", postHandler.List)
			posts.GET("/:id", postHandler.Get)
			posts.PUT("/:id", postHandler.Update)
			posts.GET("/:id", postHandler.Delete)
		}
	}

	e.Logger.Fatal(e.Start(":8080"))
}
