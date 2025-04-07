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

// gènère un token de session aléatoire
func GenerateSessionToken() string {
	bytes := make([]byte, 32)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "invalid method", http.StatusBadRequest)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	var hashedpassword string
	// Vérifie si l'utilisateur existe et récupère le mot de passe haché
	err := Database.DB.QueryRow("SELECT password FROM users WHERE username=?", username).Scan(&hashedpassword)
	if err == sql.ErrNoRows {
		json.NewEncoder(w).Encode(map[string]string{"message": "Utilisateur introuvable"})
		return
	}
	// Vérifie si le mot de passe correspond
	err = bcrypt.CompareHashAndPassword([]byte(hashedpassword), []byte(password))
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"message": "Mot de passe incorect"})
		return
	}

	token := GenerateSessionToken()
	Database.DB.Exec("UPDATE users SET token=? WHERE username=?", token, username)
	cookie := &http.Cookie{
		Name:     "session",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
	}
	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/", http.StatusSeeOther)

}
