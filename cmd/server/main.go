package main

import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.String() == "/" {
		fmt.Fprintln(w, "I love Sophie!")
	} else {
		http.NotFound(w, r)
	}
}

func cat(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./web/html/test.html")
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/sleepy", cat)

	fileServer := http.FileServer(http.Dir("./web/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	fmt.Println("Starting server...")
	http.ListenAndServe(":8090", mux)
}
