package repository

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pranavp37/go_todo_ptsql/internal/models"
)

func CreateTodorepo(connpool *pgxpool.Pool, todo *models.Todo) (*models.TodoResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var rtodo models.TodoResponse
	query := `INSERT INTO todos (title,description,users_id,created_at,updated_at) VALUES ($1,$2,$3,$4,$5) RETURNING id,title,description,users_id,created_at, updated_at`
	err := connpool.QueryRow(ctx, query, todo.Title, todo.Description, todo.UserID, todo.CreatedAt, todo.UpdatedAt).Scan(&rtodo.ID, &rtodo.Title, &rtodo.Description, &rtodo.UserID, &rtodo.CreatedAt, &rtodo.UpdatedAt)
	if err != nil {
		log.Printf("failed to create todo %v", err)
		return nil, err
	}

	return &rtodo, nil
}

func GetAllTodoRepo(connpool *pgxpool.Pool) ([]models.TodoResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	query := `SELECT id, title, description, users_id, created_at, updated_at FROM todos`
	rows, err := connpool.Query(ctx, query)
	if err != nil {
		log.Printf("Failed to fetch all todos")
		return nil, err

	}
	defer rows.Close()
	var todos []models.TodoResponse
	for rows.Next() {
		var todo models.TodoResponse
		err := rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.UserID, &todo.CreatedAt, &todo.UpdatedAt)
		if err != nil {
			log.Printf("Failed to scan todo: %v", err)
			return nil, err
		}
		todos = append(todos, todo)
	}
	return todos, nil
}

func UpdateTodoRepo(connpool *pgxpool.Pool, todo *models.Todo) (*models.TodoResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var rtodo models.TodoResponse
	query := `UPDATE todos SET title=$1, description=$2, users_id=$3, updated_at=$4 WHERE id=$5 RETURNING id,title,description,users_id,created_at, updated_at`
	err := connpool.QueryRow(ctx, query, todo.Title, todo.Description, todo.UserID, time.Now(), todo.ID).Scan(&rtodo.ID, &rtodo.Title, &rtodo.Description, &rtodo.UserID, &rtodo.CreatedAt, &rtodo.UpdatedAt)
	if err != nil {
		log.Printf("failed to update todo %v", err)
		return nil, err
	}
	return &rtodo, nil
}

func DeleteTodoRepo(connpool *pgxpool.Pool, todoID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	query := `DELETE FROM todos WHERE id=$1`
	_, err := connpool.Exec(ctx, query, todoID)
	if err != nil {
		if err.Error() == "no rows affected" {
			log.Printf("todo with id %v not found", todoID)
			return fmt.Errorf("todo not found")

		}
		log.Printf("failed to delete todo %v", err)
		return err
	}
	return nil
}
