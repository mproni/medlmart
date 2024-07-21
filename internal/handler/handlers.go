package handler

import (
	"fmt"
	"net/http"
)

func SweetHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.String() == "/" {
		fmt.Fprintln(w, "I love Sophie!")
	} else {
		http.NotFound(w, r)
	}
}

func Cat(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./web/html/test.html")
}

func Home(w http.ResponseWriter, r *http.Request) {
	// TODO
}

func Food(w http.ResponseWriter, r *http.Request) {
	// TODO
}

func Handicraft(w http.ResponseWriter, r *http.Request) {
	// TODO
}

func Exotic(w http.ResponseWriter, r *http.Request) {
	// TODO
}

func Animal(w http.ResponseWriter, r *http.Request) {
	// TODO
}

func Cart(w http.ResponseWriter, r *http.Request) {
	// TODO
}
