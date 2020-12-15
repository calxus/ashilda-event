package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

var basePath = "api/v1"

func GetEvent(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Get Request for Event!")
	w.WriteHeader(200)
}

func HandleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc(basePath + "/event", GetEvent).Methods("GET")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {
	HandleRequests()
}