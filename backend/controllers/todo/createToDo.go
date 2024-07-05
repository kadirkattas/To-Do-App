package todo

import (
	"encoding/json"
	"net/http"
	"time"

	"todo/backend/controllers/structs"
	"todo/backend/database"

	"github.com/mattn/go-sqlite3"
)

func CreateToDo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusBadRequest)
		return
	}

	var todo structs.ToDo

	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if todo.Title == "" {
		http.Error(w, "Title cannot empty", http.StatusBadRequest)
		return
	}

	v, _ := time.Now().UTC().MarshalText()

	stmt, err := database.DB.Prepare("INSERT INTO TODO (UserID, UserName, Title, Description, PostDate) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		if sqliteErr, ok := err.(sqlite3.Error); ok && sqliteErr.ExtendedCode == sqlite3.ErrConstraintUnique {
			http.Error(w, "Email already exists", http.StatusConflict)
			return
		}
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	var toDoID int
	database.DB.QueryRow("SELECT ID FROM TODO WHERE PostDate = ?", string(v)).Scan(&toDoID)

	todo.ID = toDoID

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(todo)
}
