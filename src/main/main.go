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
	log.Println("Server listening on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Hello there world")
}