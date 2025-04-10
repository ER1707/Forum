package API

import (
	"Forum/Database"
	"encoding/json"
	"log"
	"net/http"
)

type Post struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Username  string `json:"username"`
	CreatedAt string `json:"created_at"`
}

func PrintPosts(w http.ResponseWriter, r *http.Request) {
	rows, err := Database.DB.Query("SELECT posts.id, posts.title, posts.content, users.username, posts.created_at FROM posts INNER JOIN users ON posts.user_id = users.id ORDER BY posts.created_at DESC")
	if err != nil {
		http.Error(w, "Erreur lors de la récupération des posts", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var posts []Post
	for rows.Next() {
		var post Post
		err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.Username, &post.CreatedAt)
		if err != nil {
			log.Println("Erreur lors de la récupération des posts :", err)

		}
		posts = append(posts, post)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}
