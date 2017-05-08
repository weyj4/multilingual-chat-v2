package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.PathPrefix("/dist/").Handler(http.StripPrefix("/dist/", http.FileServer(http.Dir("dist/"))))
	r.HandleFunc("/", IndexHandler)

	hub := NewHub()
	go hub.run()
	r.HandleFunc("/socket", func(w http.ResponseWriter, r *http.Request) {
		HandleClient(w, r, hub)
	})

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("ListenAndServe:", err.Error())
	}
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./index.html")
}
