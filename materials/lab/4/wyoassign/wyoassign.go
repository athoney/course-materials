package wyoassign

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

type Response struct{
	Assigns []Assignment `json:"assignments"`
}

type ResponseAssign struct{
	Assign Assignment `json:"assignment"`
}

type Assignment struct {
	PK int `json:"pk"`
	Id string `json:"id"`
	Class string `json:"class"`
	Title string `json:"title"`
	Description string `json:"desc"`
	Points int `json:"points"`
	DueDate string `json:"duedate"`
	TimeEstimate string  `json:"timeestimate"`
}

var Assignments []Assignment
const Valkey string = "FooKey"
var templates = template.Must(template.ParseFiles("/home/cabox/workspace/course-materials/materials/lab/4/wyoassign/home.html", 
	"/home/cabox/workspace/course-materials/materials/lab/4/wyoassign/newAssignment.html", 
	"/home/cabox/workspace/course-materials/materials/lab/4/wyoassign/assignments.html", 
	"/home/cabox/workspace/course-materials/materials/lab/4/wyoassign/header.html", 
	"/home/cabox/workspace/course-materials/materials/lab/4/wyoassign/footer.html"))


func InitAssignments(){
	var assignmnet Assignment
	assignmnet.PK = 0
	assignmnet.Id = "3010"
	assignmnet.Class = "Software Design"
	assignmnet.Title = "Program02"
	assignmnet.Description = "Next programming assignment"
	assignmnet.Points = 100
	assignmnet.DueDate = "March 4, 2022"
	assignmnet.TimeEstimate = "Minute(s)"
	Assignments = append(Assignments, assignmnet)
}

func APISTATUS(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s end point", r.URL.Path)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "API is up and running")
}


func GetAssignments(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s end point", r.URL.Path)
	var response Response

	response.Assigns = Assignments

	//w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	
	for _, val := range Assignments {
		log.Print(val.PK)
	} 
	templates.ExecuteTemplate(w, "assign", response)
}

func GetAssignmentId(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s end point", r.URL.Path)
	//w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	params := mux.Vars(r)
	var AssignmentIds []Assignment
	var response Response

	for i, assignment := range Assignments {
		if assignment.Id == params["id"]{
			AssignmentIds = append(AssignmentIds, Assignments[i])
		}
	}
	response.Assigns = AssignmentIds
	templates.ExecuteTemplate(w, "assign", response)
}

func DeleteAssignment(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s DELETE end point", r.URL.Path)
	w.Header().Set("Content-Type", "application/txt")
	w.WriteHeader(http.StatusOK)
	params := mux.Vars(r)
	
	key := r.URL.Query().Get("validationkey")

	response := make(map[string]string)
	response["validationKey"] = Valkey

	if key == Valkey {
		response["status"] = "No Such ID to Delete"
		for index, assignment := range Assignments {
				if assignment.Id == params["id"]{
					Assignments = append(Assignments[:index], Assignments[index+1:]...)
					response["status"] = "Success"
					break
				}
		}
	}	
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return
	}
	w.Write(jsonResponse)
}

func ModifyAssignment(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s end point", r.URL.Path)
	// w.Header().Set("Content-Type", "application/json")
	if err := r.ParseForm(); err != nil {
		log.Print(w, "ParseForm() err: %v", err)
		return
	}
	var response ResponseAssign

	if (r.FormValue("modify") == "update"){
		log.Print(r.FormValue("modify"))
		//New template
		//Loop through 
		PK, _ := strconv.Atoi(r.FormValue("PK"))
		response.Assign = Assignments[PK]
		templates.ExecuteTemplate(w, "newAssign", response)
		log.Print(Assignments[PK])

	 } else {
		log.Print(r.FormValue("modify"))
	 	DeleteAssignment(w, r)
	 	templates.ExecuteTemplate(w, "newAssign", r)
	} 

}

func CreateAssignment(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "application/json")
	if err := r.ParseForm(); err != nil {
		log.Print(w, "ParseForm() err: %v", err)
		return
	}
	var assignmnet Assignment
	
	// Possible TODO: Better Error Checking!
	// Possible TODO: Better Logging
	log.Printf("In create assignment")
	log.Print(r.FormValue("class"))

	if(r.FormValue("id") != ""){
		assignmnet.PK = Assignments[len(Assignments)-1].PK + 1
		assignmnet.Id =  r.FormValue("id")
		assignmnet.Title =  r.FormValue("title")
		assignmnet.Class =  r.FormValue("class")
		assignmnet.Description =  r.FormValue("desc")
		assignmnet.Points, _ =  strconv.Atoi(r.FormValue("points"))
		assignmnet.DueDate =  r.FormValue("duedate")
		assignmnet.TimeEstimate =  r.FormValue("timeestimate")
		Assignments = append(Assignments, assignmnet)
		w.WriteHeader(http.StatusCreated)
		templates.ExecuteTemplate(w, "newAssign", struct{ Success bool}{true})
		return
	}
	w.WriteHeader(http.StatusNotFound)

}

func NewAssignment(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s end point", r.URL.Path)

	templates.ExecuteTemplate(w, "newAssign", nil)
	
	//TODO This should like like cross betweeen Create / Get   

}

func Home(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s end point", r.URL.Path)

	templates.ExecuteTemplate(w, "home", struct{ Name string}{"Alicia"})
	
	//TODO This should like like cross betweeen Create / Get   

}