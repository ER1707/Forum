package API

import (
	"Forum/Database"
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func GenerateSessionToken() string {
	bytes := make([]byte, 32)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "invalid method", http.StatusBadRequest)
	}

	username := r.FormValue("username")
	password := r.FormValue("password")
	if username == "" || password == "" {
		http.Error(w, "Missing parameter", http.StatusBadRequest)
		return
	}
	var hashedpassword string
	err := Database.DB.QueryRow("SELECT password FROM users WHERE username=?", username).Scan(&hashedpassword)
	if err == sql.ErrNoRows {
		json.NewEncoder(w).Encode(map[string]string{"message": "Utilisateur introuvable"})
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(hashedpassword), []byte(password))
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"message": "Mot de passe incorect"})
		return
	}

	token := GenerateSessionToken()
	Database.DB.Exec("UPDATE users SET tokken=? WHERE username=?", token, username)
	cookie := &http.Cookie{
		Name:     "session",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
	}
	http.SetCookie(w, cookie)
	json.NewEncoder(w).Encode(map[string]string{"message": "Utilisateur connecté"})

}
