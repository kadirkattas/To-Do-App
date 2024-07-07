package todo

import (
	"encoding/json"
	"net/http"

	"todo/backend/auth"
	"todo/backend/controllers/structs"
	"todo/backend/database"
)

func GetToDo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusBadRequest)
		return
	}

	userID := r.Context().Value(auth.ContextUserID).(int)

	rows, err := database.DB.Query("SELECT ID, Title, Description, UserID, UserName, PostDate FROM TODO WHERE UserID = ?", userID)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var todos []structs.ToDo
	for rows.Next() {
		var todo structs.ToDo
		err := rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.UserID, &todo.UserName, &todo.PostDate)
		if err != nil {
			http.Error(w, "Error scanning database result", http.StatusInternalServerError)
			return
		}
		todos = append(todos, todo)
	}

	if err = rows.Err(); err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(todos)
}
