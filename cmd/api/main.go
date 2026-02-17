package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/pranavp37/go_todo_ptsql/internal/config"
	"github.com/pranavp37/go_todo_ptsql/internal/database"
	"github.com/pranavp37/go_todo_ptsql/internal/handlers"
	"github.com/pranavp37/go_todo_ptsql/internal/middleware"
)

func main() {
	//echo initialization
	c := echo.New()

	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Failed to load configuration:", err)
	}

	connpool, err := database.Connect(cfg.DATABASE_URL)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	c.GET("/", func(c echo.Context) error {
		return c.String(200, "hello world...")
	})

	c.POST("/create", handlers.CreateUserHandeler(connpool))
	c.POST("/login", handlers.LoginUserHandeler(connpool))

	middleware.RegisterMiddleware(c)
	protected := c.Group("", middleware.AuthJwtMiddleware())

	protected.GET("/user/:id", handlers.GetUserByIdHandeler(connpool))
	protected.POST("/create-todo", handlers.CreateTodorHander(connpool))
	protected.GET("/todos", handlers.GetAllTodoHandeler(connpool))
	protected.PUT("/update-todo/:id", handlers.UpdateTodoHander(connpool))
	protected.DELETE("/delete-todo/:id", handlers.DeleteTodoHandler(connpool))

	c.Logger.Fatal(c.Start(":8080"))

}
