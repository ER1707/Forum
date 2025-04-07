package API

import (
	"Forum/Database"
	"net/http"
)

type Post struct {
	UserID     int    `json:"user_id"`
	CategoryID int    `json:"category_id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
}

func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	category := r.FormValue("category")
	title := r.FormValue("title")
	content := r.FormValue("content")
	if category == "" || title == "" || content == "" {
		http.Error(w, "Tous les champs sont obligatoires", http.StatusBadRequest)
		return
	}

	user, uerr := Database.UserInfos(w, r)
	if !uerr {
		http.Error(w, "Erreur lors de la création du post", http.StatusInternalServerError)
		return
	}
	_, err := Database.DB.Exec("INSERT INTO posts (user_id, category_id, title, content) VALUES (?, ?, ?, ?)", user.ID, category, title, content)
	if err != nil {
		http.Error(w, "Erreur lors de la création du post", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Post créé avec succès !"))
}
