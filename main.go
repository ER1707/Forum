package main

import (
	"fmt"
	"net/http"
	"log"
	"html/template"
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

func main() {
	// Définir le handler pour la route /
	http.HandleFunc("/", homeHandler)

	// Définir le port du serveur
	port := ":8080"
	fmt.Println("Serveur démarré sur http://localhost" + port)

	// Démarrer le serveur
	http.ListenAndServe(port, nil)
}
