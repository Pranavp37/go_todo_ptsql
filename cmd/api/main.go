package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/pranavp37/go_todo_ptsql/internal/config"
	"github.com/pranavp37/go_todo_ptsql/internal/database"
)

func main() {
	//echo initialization
	c := echo.New()

	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Failed to load configuration:", err)
	}

	_, err = database.Connect(cfg.DATABASE_URL)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	c.GET("/", func(c echo.Context) error {
		return c.String(200, "hello world...")
	})

	c.Start(":8080")

}
