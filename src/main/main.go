package main

import (
    "encoding/json"
    "log"
    "net/http"
	"github.com/gorilla/mux"
)

func main(){
	router := mux.NewRouter()

	router.HandleFunc("/", HelloWorld).Methods("GET")
	router.HandleFunc("/hello", HelloWorld).Methods("GET")
	log.Println("Server listening on port 18220")
	log.Fatal(http.ListenAndServe(":18220", router))
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Hello world")
}