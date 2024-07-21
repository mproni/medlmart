package server

import (
	"fmt"
	"net/http"

	"github.com/mproni/medlmart/internal/handler"
)

func Start() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handler.SweetHome)
	mux.HandleFunc("/sleepy", handler.Cat)

	fileServer := http.FileServer(http.Dir("./web/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	fmt.Println("Starting server...")
	http.ListenAndServe(":8090", mux)
}
