package API

import (
	"Forum/Database"
	"encoding/json"
	"net/http"
)

type Post struct {
	UserID     int    `json:"user_id"`
	CategoryID int    `json:"category_id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
}

func CreatePost(userID int, categoryID int, title string, content string) error {
	_, err := Database.DB.Exec(`INSERT INTO posts (user_id, category_id, title, content) VALUES (?, ?, ?, ?)`, userID, categoryID, title, content)
	return err
}

func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	var post Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		http.Error(w, "Données invalides", http.StatusBadRequest)
		return
	}

	err = CreatePost(post.UserID, post.CategoryID, post.Title, post.Content)
	if err != nil {
		http.Error(w, "Erreur lors de la création du post", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Post créé avec succès !"))
}
