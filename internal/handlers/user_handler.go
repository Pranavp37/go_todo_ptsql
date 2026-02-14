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

func CreateUserHandeler(connpool *pgxpool.Pool) echo.HandlerFunc {

	return func(c echo.Context) error {
		var user models.User
		if err := c.Bind(&user); err != nil {
			return c.JSON(http.StatusBadRequest, utiles.Response{
				Success: false,
				Message: "Invalid input",
			})
		}

		var userModel = &models.User{
			ID:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			Password:  user.Password,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		}

		err := repository.CreateUser(connpool, userModel)
		if err != nil {
			if err.Error() == "user already exists" {
				return c.JSON(http.StatusConflict, utiles.Response{
					Success: false,
					Message: "User already exists",
				})
			}
			return c.JSON(http.StatusInternalServerError, utiles.Response{
				Success: false,
				Message: "Failed to create user",
			})
		}

		return c.JSON(http.StatusCreated, utiles.Response{
			Success: true,
			Message: "User created successfully",
		})
	}
}

func LoginUserHandeler(connpool *pgxpool.Pool) echo.HandlerFunc {
	return func(c echo.Context) error {
		var user models.User
		if err := c.Bind(&user); err != nil {
			log.Print("Error binding user input: ", err)
			return c.JSON(http.StatusBadRequest, utiles.Response{
				Success: false,
				Message: "Invalid input",
			})

		}
		usermodel, err := repository.LoginUser(connpool, &user)
		if err != nil {
			log.Print("Error logging in user: ", err)
			return c.JSON(http.StatusUnauthorized, utiles.Response{
				Success: false,
				Message: "Invalid email or password",
			})
		}

		return c.JSON(http.StatusOK, utiles.Response{
			Success: true,
			Message: "User logged in successfully",
			Data:    usermodel,
		})
	}

}
