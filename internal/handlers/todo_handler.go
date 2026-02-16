package handlers

import (
	"log"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"

	"github.com/pranavp37/go_todo_ptsql/internal/models"
	"github.com/pranavp37/go_todo_ptsql/internal/repository"
	"github.com/pranavp37/go_todo_ptsql/internal/utiles"
)

func CreateTodorHander(connpool *pgxpool.Pool) echo.HandlerFunc {
	return func(c echo.Context) error {

		var todoreq models.Todo
		if err := c.Bind(&todoreq); err != nil {
			log.Print("Error binding user input: ", err)
			return c.JSON(http.StatusBadRequest, utiles.Response{
				Success: false,
				Message: "Invalid input",
			})
		}

		todo, err := repository.CreateTodorepo(connpool, &todoreq)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, utiles.Response{
				Success: false,
				Message: "create todo failed",
			})
		}

		return c.JSON(http.StatusCreated, utiles.Response{
			Success: true,
			Message: "todo created successfully",
			Data:    todo,
		})

	}
}

func GetAllTodoHandeler(connpool *pgxpool.Pool) echo.HandlerFunc {
	return func(c echo.Context) error {
		todos, err := repository.GetAllTodoRepo(connpool)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, utiles.Response{
				Success: false,
				Message: "Failed to fetch todos",
			})
		}

		return c.JSON(http.StatusOK, utiles.Response{
			Success: true,
			Message: "Todos fetched successfully",
			Data:    todos,
		})
	}
}

func UpdateTodoHander(connpool *pgxpool.Pool) echo.HandlerFunc {
	return func(c echo.Context) error {
		var todo models.Todo
		if err := c.Bind(&todo); err != nil {
			log.Print("")
			log.Print("Error binding user input: ", err)
			return c.JSON(http.StatusBadRequest, utiles.Response{
				Success: false,
				Message: "Invalid input",
			})
		}
		rtodo, err := repository.UpdateTodoRepo(connpool, &todo)
		if err != nil {
			log.Print("updating todo table failed")
			return c.JSON(http.StatusInternalServerError, utiles.Response{
				Success: false,
				Message: "failed to update todo",
				Data:    err,
			})
		}

		return c.JSON(http.StatusOK, utiles.Response{
			Success: true,
			Message: "Todo updated successfully",
			Data:    rtodo,
		})

	}

}

func DeleteTodoHandler(connpool *pgxpool.Pool) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		if err := repository.DeleteTodoRepo(connpool, id); err != nil {
			if err.Error() == "todo not found" {
				return c.JSON(http.StatusNotFound, utiles.Response{
					Success: false,
					Message: "Todo not found",
				})
			}
			log.Printf("Failed to detete table")
			return c.JSON(http.StatusInternalServerError, utiles.Response{
				Success: false,
				Message: "Failed to Detete Data",
				Data:    err,
			})
		}
		return c.JSON(http.StatusOK, utiles.Response{
			Success: true,
			Message: "Todo deleted successfully",
		})

	}
}
