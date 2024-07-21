package handler

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func SweetHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.String() == "/sweet" {
		fmt.Fprintln(w, "I love Sophie!")
	} else {
		http.NotFound(w, r)
	}
}
/*
func Cat(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./web/html/test.html")
}
*/
func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.String() != "/" {
		http.NotFound(w, r)
		return
	}

	files := []string{
		"./web/html/home.page.tmpl",
		"./web/html/base.layout.tmpl",
		"./web/html/footer.partial.tmpl",
	}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
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
