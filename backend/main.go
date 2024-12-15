package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	
	"github.com/gorilla/mux"
	"github.com/lib/pq"
)

type User struct{
	Id		int 	`json:"id"`
	Name 	string	`json:"name"`
	Email 	string	`json:"email"`
}


func main(){
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	
	defer db.Close()
	

	// create db if not exists
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, name TEXT, email TEXT)")
	if err != nil{
		log.Fatal(err)
	}
	
	//router
	router := mux.NewRouter()
	router.HandleFunc("/api/go/users", getUsers(db)).Methods("GET")
	router.HandleFunc("/api/go/users", createUser(db)).Methods("POST")
	router.HandleFunc("/api/go/users/{id}", getUser(db)).Methods("GET")
	router.HandleFunc("/api/go/users/{id}", updateUser(db)).Methods("PUT")
	router.HandleFunc("/api/go/users/{id}", deleteUser(db)).Methods("DELETE")
	
	enhancedRouter := enableCORS(jsonContentTypeMiddleware(router))

	log.Fatal(http.ListenAndServe(":8000", enhancedRouter))
}