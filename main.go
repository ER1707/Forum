package main

import (
	"Forum/API"
	"Forum/Database"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// Handler pour la page d'accueil
func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Println(err)
		return
	}
	tmpl.Execute(w, nil)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/login.html")
	if err != nil {
		log.Println("Erreur lors du chargement de la page de connexion :", err)
		return
	}
	tmpl.Execute(w, nil)
}

func main() {
	err := Database.InitDB()
	if err != nil {
		log.Fatal("Erreur lors de l'initialisation de la base de données :", err)
	}
	defer Database.CloseDB()
	// Définir le handler pour la route /

	http.Handle("/styles/", http.StripPrefix("/styles/", http.FileServer(http.Dir("styles"))))
	http.Handle("/scripts/", http.StripPrefix("/scripts/", http.FileServer(http.Dir("scripts"))))
	http.HandleFunc("/login", loginHandler)

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/api/register", API.Register)
	http.HandleFunc("/api/login", API.Login)

	// Définir le port du serveur
	port := ":8080"
	fmt.Println("Serveur démarré sur http://localhost" + port)

	// Démarrer le serveur
	http.ListenAndServe(port, nil)
}
