package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

func postHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("post request")

	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)

	if err != nil {
		fmt.Println("Error reading")
		return
	}

	fmt.Println(string(body))

}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/hello", postHandle).Methods("POST")
	fmt.Println("Servidors escuchando en el puertoss 8080")

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println(err.Error())
	}
}
