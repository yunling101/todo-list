package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"github.com/google/uuid"
)

type Todo struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type TodoInput struct {
	Title     string  `json:"title"`
	Completed *bool   `json:"completed"`
}

var db *sql.DB

func initDB() error {
	var err error
	db, err = sql.Open("sqlite3", "./todos.db")
	if err != nil {
		return err
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS todos (
			id TEXT PRIMARY KEY,
			title TEXT NOT NULL,
			completed INTEGER DEFAULT 0,
			created_at TEXT NOT NULL,
			updated_at TEXT NOT NULL
		)
	`)
	return err
}

func main() {
	if err := initDB(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	r := gin.Default()

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// API routes
	v1 := r.Group("/api/v1")
	{
		v1.GET("/todos", getTodos)
		v1.POST("/todos", createTodo)
		v1.GET("/todos/:id", getTodo)
		v1.PUT("/todos/:id", updateTodo)
		v1.DELETE("/todos/:id", deleteTodo)
		v1.PATCH("/todos/:id/toggle", toggleTodo)
	}

	log.Println("Server starting on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func getTodos(c *gin.Context) {
	rows, err := db.Query("SELECT id, title, completed, created_at, updated_at FROM todos ORDER BY created_at DESC")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var todos []Todo
	for rows.Next() {
		var todo Todo
		var createdAt, updatedAt string
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Completed, &createdAt, &updatedAt); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		todo.CreatedAt, _ = time.Parse(time.RFC3339, createdAt)
		todo.UpdatedAt, _ = time.Parse(time.RFC3339, updatedAt)
		todos = append(todos, todo)
	}

	if todos == nil {
		todos = []Todo{}
	}
	c.JSON(http.StatusOK, todos)
}

func createTodo(c *gin.Context) {
	var input TodoInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	now := time.Now().UTC()
	todo := Todo{
		ID:        uuid.New().String(),
		Title:     input.Title,
		Completed: false,
		CreatedAt: now,
		UpdatedAt: now,
	}

	_, err := db.Exec(
		"INSERT INTO todos (id, title, completed, created_at, updated_at) VALUES (?, ?, ?, ?, ?)",
		todo.ID, todo.Title, todo.Completed, todo.CreatedAt.Format(time.RFC3339), todo.UpdatedAt.Format(time.RFC3339),
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, todo)
}

func getTodo(c *gin.Context) {
	id := c.Param("id")

	var todo Todo
	var createdAt, updatedAt string
	err := db.QueryRow("SELECT id, title, completed, created_at, updated_at FROM todos WHERE id = ?", id).
		Scan(&todo.ID, &todo.Title, &todo.Completed, &createdAt, &updatedAt)

	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "todo not found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	todo.CreatedAt, _ = time.Parse(time.RFC3339, createdAt)
	todo.UpdatedAt, _ = time.Parse(time.RFC3339, updatedAt)
	c.JSON(http.StatusOK, todo)
}

func updateTodo(c *gin.Context) {
	id := c.Param("id")

	var input TodoInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	now := time.Now().UTC()

	var query string
	var args []interface{}

	if input.Completed != nil {
		// Update both title and completed
		query = "UPDATE todos SET title = ?, completed = ?, updated_at = ? WHERE id = ?"
		args = []interface{}{input.Title, *input.Completed, now.Format(time.RFC3339), id}
	} else {
		// Only update title, keep completed unchanged
		query = "UPDATE todos SET title = ?, updated_at = ? WHERE id = ?"
		args = []interface{}{input.Title, now.Format(time.RFC3339), id}
	}

	result, err := db.Exec(query, args...)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "todo not found"})
		return
	}

	// Fetch updated todo
	var todo Todo
	var createdAt, updatedAt string
	err = db.QueryRow("SELECT id, title, completed, created_at, updated_at FROM todos WHERE id = ?", id).
		Scan(&todo.ID, &todo.Title, &todo.Completed, &createdAt, &updatedAt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	todo.CreatedAt, _ = time.Parse(time.RFC3339, createdAt)
	todo.UpdatedAt, _ = time.Parse(time.RFC3339, updatedAt)
	c.JSON(http.StatusOK, todo)
}

func deleteTodo(c *gin.Context) {
	id := c.Param("id")

	result, err := db.Exec("DELETE FROM todos WHERE id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "todo not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "todo deleted"})
}

func toggleTodo(c *gin.Context) {
	id := c.Param("id")

	// Get current state
	var completed bool
	err := db.QueryRow("SELECT completed FROM todos WHERE id = ?", id).Scan(&completed)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "todo not found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Toggle
	now := time.Now().UTC()
	newCompleted := !completed
	result, err := db.Exec(
		"UPDATE todos SET completed = ?, updated_at = ? WHERE id = ?",
		newCompleted, now.Format(time.RFC3339), id,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "todo not found"})
		return
	}

	// Fetch updated todo
	var todo Todo
	var createdAt, updatedAt string
	err = db.QueryRow("SELECT id, title, completed, created_at, updated_at FROM todos WHERE id = ?", id).
		Scan(&todo.ID, &todo.Title, &todo.Completed, &createdAt, &updatedAt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	todo.CreatedAt, _ = time.Parse(time.RFC3339, createdAt)
	todo.UpdatedAt, _ = time.Parse(time.RFC3339, updatedAt)
	c.JSON(http.StatusOK, todo)
}
