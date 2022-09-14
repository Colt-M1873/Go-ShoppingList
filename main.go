package main

import (
	"Go-ShoppingList/api"
	"net/http"
)

func main() {
	// http.HandleFunc("/hello-world", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("hello world from go mini server"))
	// })
	// http.ListenAndServe(":8080", nil)
	srv := api.NewServer()
	http.ListenAndServe(":8080", srv)

}
