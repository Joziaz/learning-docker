package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

func getTodosHandler(w http.ResponseWriter, _ *http.Request, db *sql.DB) {
	const query = `SELECT id, title, completed FROM todos`
	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, "Failed to retrieve todos", http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	var todos []Todo
	for rows.Next() {
		var todo Todo
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Completed); err != nil {
			http.Error(w, "Failed to scan todo", http.StatusInternalServerError)
			return
		}
		todos = append(todos, todo)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(todos)
}

func createTodoHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	var todo Todo
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	const insertQuery = `INSERT INTO todos (title, completed) VALUES ($1, $2) RETURNING id`
	result := db.QueryRow(insertQuery, todo.Title, todo.Completed)

	var id int
	if err := result.Scan(&id); err != nil {
		http.Error(w, "Failed to retrieve todo ID", http.StatusInternalServerError)
		return
	}
	todo.ID = id

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(todo)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Docker!"))
}
