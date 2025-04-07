package API

import (
	"Forum/Database"
	"encoding/json"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	username := r.FormValue("username")
	password := r.FormValue("password")
	email := r.FormValue("email")

	password = Database.HashPassword(password)
	_, err := Database.DB.Exec("INSERT INTO users (username, password, email) VALUES (?, ?, ?)", username, password, email)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"message": "Pseudo ou email déjà utilisé"})
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"message": "Succès"})
}
