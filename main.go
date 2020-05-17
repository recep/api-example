package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type computer struct {
	ID    string `json:"id"`
	Brand string `json:"brand"`
	Model string `json:"model"`
	Price int    `json:"price`
}

var computers []computer

func main() {

	computers = []computer{
		{ID: "1", Brand: "APPLE", Model: "Pro", Price: 10000},
		{ID: "2", Brand: "DELL", Model: "XPS", Price: 9000},
	}

	handleRequest()
}

func handleRequest() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", index)
	router.HandleFunc("/computers", returnAllComputers).Methods("GET")
	router.HandleFunc("/computer/{id}", returnSingleComputer).Methods("GET")
	router.HandleFunc("/computer", createNewComputer).Methods("POST")

	log.Fatal(http.ListenAndServe(":8081", router))
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "İndex Page")
}

func returnAllComputers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(computers)
}


func returnSingleComputer(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	for _, computer := range computers {
		if computer.ID == id {
			json.NewEncoder(w).Encode(computer)
		}
	}
}

func createNewComputer(w http.ResponseWriter, r *http.Request) {
	var computer computer

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(reqBody, &computer)
	computers = append(computers, computer)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(computer)
}


