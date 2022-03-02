package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"wyoassign/wyoassign"

)


func main() {
	wyoassign.InitAssignments() //creates baseline data
	log.Println("starting API server")
	//create a new router
	router := mux.NewRouter()
	log.Println("creating routes")
	//specify endpoints
	router.HandleFunc("/api-status", wyoassign.APISTATUS).Methods("GET")
	router.HandleFunc("/assignments", wyoassign.GetAssignments).Methods("GET")
	router.HandleFunc("/assignment/{id}", wyoassign.GetAssignmentId).Methods("GET")
	router.HandleFunc("/assignment/{id}", wyoassign.DeleteAssignmentAPI).Queries("validationkey", "{key}").Methods("DELETE")		
	router.HandleFunc("/newAssignment", wyoassign.NewAssignment).Methods("GET")
	router.HandleFunc("/newAssignment", wyoassign.CreateAssignment).Methods("POST")
	router.HandleFunc("/assignment/modify", wyoassign.ModifyAssignment).Methods("POST")
	router.HandleFunc("/assignment/modify/update", wyoassign.UpdateAssignment).Methods("POST")
	router.HandleFunc("/assignment/modify/delete", wyoassign.DeleteAssignment).Methods("POST")

	
	router.HandleFunc("/", wyoassign.Home)

	// router.HandleFunc("/assignments/{id}", wyoassign.UpdateAssignment).Methods("PUT")

	http.Handle("/", router)

	//start and listen to requests
	http.ListenAndServe(":8080", router)

}