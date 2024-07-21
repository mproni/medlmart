package server

import (
	"fmt"
	"net/http"

	"github.com/mproni/medlmart/internal/handler"
)

func StartServer() {
	mux := http.NewServeMux()

	mux.HandleFunc("/sweet", handler.SweetHome)
//	mux.HandleFunc("/sleepy", handler.Cat)

	mux.HandleFunc("/", handler.Home)
	mux.HandleFunc("/food", handler.Food)
	mux.HandleFunc("/handicraft", handler.Handicraft)
	mux.HandleFunc("/exotic", handler.Exotic)
	mux.HandleFunc("/animal", handler.Animal)
	mux.HandleFunc("/cart", handler.Cart)

	fileServer := http.FileServer(http.Dir("./web/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	fmt.Println("Starting server...")
	http.ListenAndServe(":8090", mux)
}
