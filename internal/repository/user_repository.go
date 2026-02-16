package repository

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pranavp37/go_todo_ptsql/internal/models"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(connpool *pgxpool.Pool, user *models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	isExisting, err := IsEmailRegistered(connpool, user.Email)
	if err != nil {
		log.Print("Error checking if user exists: ", err)
		return err
	}
	if isExisting {
		log.Print("User already exists: ", user.Email)
		return fmt.Errorf("user already exists")
	}

	hashpass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		log.Print("Error hashing password: ", err)
		return err
	}

	query := `INSERT INTO users (name,email,password) VALUES ($1,$2,$3)`

	_, err = connpool.Exec(ctx, query, user.Name, user.Email, hashpass)
	if err != nil {
		log.Print("Error creating user: ", err)
		return err
	}
	return nil
}

func IsEmailRegistered(connpool *pgxpool.Pool, email string) (bool, error) {
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

func LoginUser(connpool *pgxpool.Pool, user *models.User) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var userModel models.User
	query := `SELECT id, name, email, password, created_at, updated_at FROM users WHERE email = $1`
	row := connpool.QueryRow(ctx, query, user.Email)
	err := row.Scan(&userModel.ID, &userModel.Name, &userModel.Email, &userModel.Password, &userModel.CreatedAt, &userModel.UpdatedAt)
	if err != nil {
		if err.Error() == "no rows in result set" {
			return nil, fmt.Errorf("invalid email or password")
		}
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(userModel.Password), []byte(user.Password))
	if err != nil {
		return nil, fmt.Errorf("password does not match")
	}
	var userResponse *models.User = &models.User{
		ID:        userModel.ID,
		Name:      userModel.Name,
		Email:     userModel.Email,
		CreatedAt: userModel.CreatedAt,
		UpdatedAt: userModel.UpdatedAt,
	}
	return userResponse, nil
}

func GetUserByID(connpool *pgxpool.Pool, userID string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var userModel models.User
	query := `SELECT id,email,name,created_at FROM users WHERE id = $1`
	row := connpool.QueryRow(ctx, query, userID)

	if err := row.Scan(&userModel.ID, &userModel.Email, &userModel.Name, &userModel.CreatedAt); err != nil {
		if err.Error() == "no rows in result set" {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err

	}
	return &userModel, nil

}
