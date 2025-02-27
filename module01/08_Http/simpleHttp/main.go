package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World, you accessed : %s", r.URL.Path)
}

func main() {
	http.HandleFunc("/", handler) //Rota
	fmt.Println("Server running in http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
