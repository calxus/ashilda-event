package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var basePath = "api/v1"
var databaseUsername = os.Getenv("DATABASE_USERNAME")
var databasePassword = os.Getenv("DATABASE_PASSWORD")
var databaseHost = os.Getenv("DATABASE_HOST")
var databasePort = os.Getenv("DATABASE_PORT")
var databaseName = os.Getenv("DATABASE_NAME")
var db sql.Conn

func getEvent(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Get Request for Event!")
	w.WriteHeader(200)
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc(basePath+"/event", getEvent).Methods("GET")
	log.Fatal(http.ListenAndServe(":8082", myRouter))
}

func main() {
	db, err := sql.Open("mysql", databaseUsername+":"+databasePassword+"@tcp("+databaseHost+":"+databasePort+")/"+databaseName)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Connection to database successful")
	defer db.Close()
	handleRequests()
}
