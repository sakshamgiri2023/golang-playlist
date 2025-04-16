package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Mod module in golang")
	greeter()

	r := mux.NewRouter()
	r.HandleFunc("/", serverHome).Methods("GET")

	http.ListenAndServe(":4000", r)

}

func greeter() {
	fmt.Println("hey there")
}

func serverHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>welcome to golang landing page</h1>"))

}
