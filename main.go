package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/hello" {
		http.NotFound(w, r)
		return
	}

	name := r.URL.Query().Get("name")

	if name == "" {
		name = "inconnu"
	}

	fmt.Fprintf(w, "Bonjour %s", name)
}

func main() {
	http.HandleFunc("/", handler)

	fmt.Println("Serveur démarré sur http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}