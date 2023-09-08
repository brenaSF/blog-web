package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "static/style.css" {
			w.Header().Set("Content-Type", "text/css")
	
			http.ServeFile(w, r, "static/home.css")
			return
		}
		http.ServeFile(w, r, "views/home.html")
	})
	

	fmt.Println("Servindo na porta 8080...")
	http.ListenAndServe(":8080", nil)
}

