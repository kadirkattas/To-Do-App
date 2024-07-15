package loginpage

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	frontendstructs "todo/frontend/frontendStructs"
	"todo/frontend/requests"
)

const loginApiUrl = "http://localhost:8080/api/v1/user/login"

func LoginPageHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "frontend/static/loginPage/login.html")
	case "POST":
		email := r.FormValue("email")
		password := r.FormValue("password")

		user := frontendstructs.User{
			Email:    email,
			Password: password,
		}

		jsonData, err := json.Marshal(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response, errR := requests.PostRequest(loginApiUrl, jsonData)
		if errR != nil {
			fmt.Println(errR)
			http.Error(w, errR.Error(), http.StatusInternalServerError)
			return
		}

		if response.StatusCode != http.StatusOK {
			http.Error(w, "Error: Internal server error", http.StatusInternalServerError)
			return
		}

		defer response.Body.Close()
		body, err := io.ReadAll(response.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var loginResponse frontendstructs.LoginResponse
		err = json.Unmarshal(body, &loginResponse)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		token := loginResponse.Token

		http.SetCookie(w, &http.Cookie{
			Name:     "token",
			Value:    token,
			HttpOnly: true,
			Path:     "/",
		})
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
