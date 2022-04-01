package main

// main.go HAS FOUR TODOS - TODO_1 - TODO_4

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"scrape/scrape"
)


const LOG_LEVEL = 0

func main() {
	
	scrape.LOG_LEVEL = LOG_LEVEL
	if LOG_LEVEL >= 1 {
		log.Println("starting API server")
	}
	//create a new router
	router := mux.NewRouter()
	if LOG_LEVEL >= 1 {
		log.Println("creating routes")
	}
	//specify endpoints
	router.HandleFunc("/", scrape.MainPage).Methods("GET")

	router.HandleFunc("/api-status", scrape.APISTATUS).Methods("GET")

	router.HandleFunc("/indexer", scrape.IndexFiles).Methods("GET")
	router.HandleFunc("/search", scrape.FindFile).Methods("GET")		
    //TODO_2 router.HandleFunc("/addsearch/{regex}", scrape.TODOREPLACE).Methods("GET")
    //TODO_3 router.HandleFunc("/clear", scrape.TODOREPLACE).Methods("GET")
    //TODO_4 router.HandleFunc("/reset", scrape.TODOREPLACE).Methods("GET")



	http.Handle("/", router)

	//start and listen to requests
	http.ListenAndServe(":8080", router)

}