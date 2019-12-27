package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "postgres"
)

type event struct {
	Name string
}
type message struct {
	Message string
}

// CREATE
func create(w http.ResponseWriter, r *http.Request) {
	var createEvent event
	reqBody, err := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &createEvent)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	sqlStatement := `INSERT INTO data (name) VALUES ($1)`
	_, err = db.Exec(sqlStatement, createEvent.Name)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	w.WriteHeader(http.StatusCreated)
	var response message
	response.Message = "Success"
	json.NewEncoder(w).Encode(response)
}

// READ
func read(w http.ResponseWriter, r *http.Request) {
	name := "None"
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	sqlStatement := `SELECT *  FROM data `
	err = db.QueryRow(sqlStatement).Scan(&name)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	var response message
	response.Message = name
	json.NewEncoder(w).Encode(response)

}

// UPDATE
func update(w http.ResponseWriter, r *http.Request) {
	var createEvent event
	reqBody, err := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &createEvent)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	sqlStatement := `UPDATE data SET name = ($1)`
	_, err = db.Exec(sqlStatement, createEvent.Name)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	var response message
	response.Message = "Success"
	json.NewEncoder(w).Encode(response)

}

// DELETE
func delete(w http.ResponseWriter, r *http.Request) {
	var createEvent event
	reqBody, err := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &createEvent)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	sqlStatement := `DELETE FROM data WHERE name = ($1)`
	_, err = db.Exec(sqlStatement, createEvent.Name)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	var response message
	response.Message = "Success"
	json.NewEncoder(w).Encode(response)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/create", create).Methods("POST")
	router.HandleFunc("/read", read).Methods("GET")
	router.HandleFunc("/update", update).Methods("PUT")
	router.HandleFunc("/delete", delete).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}
