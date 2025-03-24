package main

import (
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
		log.Println(err)
		return
	}
	tmpl.Execute(w, nil)
}

func main() {
	// Définir le handler pour la route /

	http.Handle("/styles/", http.StripPrefix("/styles/", http.FileServer(http.Dir("styles"))))
	http.Handle("/scripts/", http.StripPrefix("/scripts/", http.FileServer(http.Dir("scripts"))))
	http.HandleFunc("/login", loginHandler)

	http.HandleFunc("/", homeHandler)

	// Définir le port du serveur
	port := ":8080"
	fmt.Println("Serveur démarré sur http://localhost" + port)

	// Démarrer le serveur
	http.ListenAndServe(port, nil)
}
