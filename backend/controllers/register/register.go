package registercontroller

import (
	"encoding/json"
	"net/http"

	"todo/backend/controllers/structs"
	"todo/backend/database"

	"github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusBadRequest)
		return
	}

	var user structs.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}

	stmt, err := database.DB.Prepare("INSERT INTO USERS (Email, UserName, Password) VALUES (?, ?, ?)")
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Email, user.UserName, string(hashedPassword))
	if err != nil {
		if sqliteErr, ok := err.(sqlite3.Error); ok && sqliteErr.ExtendedCode == sqlite3.ErrConstraintUnique {
			http.Error(w, "Email already exists", http.StatusConflict)
			return
		}
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	var UserID int
	database.DB.QueryRow("SELECT ID FROM USERS WHERE Email = ?", user.Email).Scan(&UserID)

	user.ID = UserID
	user.Password = string(hashedPassword)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
