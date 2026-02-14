package repository

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pranavp37/go_todo_ptsql/internal/models"
)

func CreateUser(connpool *pgxpool.Pool, user *models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	isExisting, err := GetUserByEmail(connpool, user.Email)
	if err != nil {
		log.Print("Error checking if user exists: ", err)
		return err
	}
	if isExisting {
		log.Print("User already exists: ", user.Email)
		return fmt.Errorf("user already exists")
	}
	
	query := `INSERT INTO users (id,name,email,password,created_at,updated_at) VALUES ($1,$2,$3,$4,$5,$6)`

	_, err = connpool.Exec(ctx, query, user.ID, user.Name, user.Email, user.Password, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		log.Print("Error creating user: ", err)
		return err
	}
	return nil
}

func GetUserByEmail(connpool *pgxpool.Pool, email string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	query := `SELECT email FROM users WHERE email = $1`
	row := connpool.QueryRow(ctx, query, email)
	var existingEmail string
	err := row.Scan(&existingEmail)
	if err != nil {
		if err.Error() == "no rows in result set" {
			return false, nil
		}
		return false, err
	}
	return true, nil
}