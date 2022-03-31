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
	Empty bool `json:"empty"`
}

type APIResponse struct{
	Assigns []Assignment `json:"assignments"`
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

type PageData struct {
	Added bool `json:"added"`
	Update bool `json:"update"`
	FormData Assignment
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

//Web-based routes

func GetAssignments(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s end point", r.URL.Path)
	var response Response

	response.Assigns = Assignments
	if (len(Assignments) == 0) {
		response.Empty = true
	} else {
		response.Empty = false
	}
	//w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	
	for _, val := range Assignments {
		log.Print(val.PK)
	}
	log.Print(response) 
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
	response.Empty = false
	templates.ExecuteTemplate(w, "assign", response)
}


func DeleteAssignmentAPI(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s DELETE end point", r.URL.Path)
	//w.Header().Set("Content-Type", "application/txt")
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
	var response PageData

	if (r.FormValue("modify") == "update"){
		log.Print(r.FormValue("modify"))
		//New template
		//Loop through 
		PK, _ := strconv.Atoi(r.FormValue("PK"))
		response.Added = false
		response.Update = true
		response.FormData = Assignments[PK]
		templates.ExecuteTemplate(w, "newAssign", response)
	 } else {
		log.Print(r.FormValue("modify"))
	 	DeleteAssignment(w, r)
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
		templates.ExecuteTemplate(w, "newAssign", struct{ Added, Update bool}{true, false})
		return
	}
	w.WriteHeader(http.StatusNotFound)

}

func UpdateAssignment(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "application/json")
	if err := r.ParseForm(); err != nil {
		log.Print(w, "ParseForm() err: %v", err)
		return
	}

	PK, _ := strconv.Atoi(r.FormValue("PK"))

	Assignments[PK].Id =  r.FormValue("id")
	Assignments[PK].Title =  r.FormValue("title")
	Assignments[PK].Class =  r.FormValue("class")
	Assignments[PK].Description =  r.FormValue("desc")
	Assignments[PK].Points, _ =  strconv.Atoi(r.FormValue("points"))
	Assignments[PK].DueDate =  r.FormValue("duedate")
	Assignments[PK].TimeEstimate =  r.FormValue("timeestimate")
	templates.ExecuteTemplate(w, "newAssign", struct{ Added, Update bool}{true, true})

}

func DeleteAssignment(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Print(w, "ParseForm() err: %v", err)
		return
	}

	for i, assignment := range Assignments {
		assignment.PK = i
	}

	PK, _ := strconv.Atoi(r.FormValue("PK"))

	if (len(Assignments) == 1) {
		Assignments = nil
	} else {
		Assignments = append(Assignments[:PK], Assignments[PK+1:]...)
	}
	var response Response
	response.Assigns = Assignments
	if (len(Assignments) == 0) {
		response.Empty = true
	} else {
		response.Empty = false
	}

	templates.ExecuteTemplate(w, "assign", response)
}

func NewAssignment(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s end point", r.URL.Path)
	templates.ExecuteTemplate(w, "newAssign", nil)
}

func Home(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s end point", r.URL.Path)
	templates.ExecuteTemplate(w, "home", struct{ Name string; NumAssigns int}{"Alicia", len(Assignments)}) 
}


//API routes
func APISTATUS(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s end point", r.URL.Path)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "API is up and running")
}

func APIGetAssignment(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s end point", r.URL.Path)
	w.Header().Set("Content-Type", "application/json")

	if err := r.ParseForm(); err != nil {
		log.Print(w, "ParseForm() err: %v", err)
		return
	}

	PK, _ := strconv.Atoi(r.FormValue("PK"))

	for _, assignment := range Assignments {
		if assignment.PK == PK{
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(assignment)
			break
		}
	}
	w.WriteHeader(http.StatusBadRequest)
}

func APIGetAssignmentId(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s end point", r.URL.Path)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	params := mux.Vars(r)

	var AssignmentsById []Assignment
	var response APIResponse

	for _, assignment := range Assignments {
		if assignment.Id == params["id"]{
			AssignmentsById = append(AssignmentsById, assignment)
		}
	}
	response.Assigns = AssignmentsById
	jsonResponse, err := json.Marshal(response)

	if err != nil {
		return
	}

	//TODO 
	w.Write(jsonResponse)
}

func APIGetAssignments(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s end point", r.URL.Path)
	var response APIResponse

	response.Assigns = Assignments

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	
	jsonResponse, err := json.Marshal(response)

	if err != nil {
		return
	}

	//TODO 
	w.Write(jsonResponse)
}

func APIAddAssignment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var assignment Assignment
	r.ParseForm()
	// Possible TODO: Better Error Checking!
	// Possible TODO: Better Logging
	if(r.FormValue("id") != ""){
		assignment.PK = Assignments[len(Assignments)-1].PK + 1
		assignment.Id =  r.FormValue("id")
		assignment.Title =  r.FormValue("title")
		assignment.Class =  r.FormValue("class")
		assignment.Description =  r.FormValue("desc")
		assignment.Points, _ =  strconv.Atoi(r.FormValue("points"))
		assignment.DueDate =  r.FormValue("duedate")
		assignment.TimeEstimate =  r.FormValue("timeestimate")
		Assignments = append(Assignments, assignment)
		w.WriteHeader(http.StatusCreated)
		return
	}
	w.WriteHeader(http.StatusNotFound)

}

func APIUpdateAssignment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	r.ParseForm()
	PK, _ := strconv.Atoi(r.FormValue("PK"))
	var idExists bool

	for _, assignment := range Assignments {
		if (assignment.PK == PK) {
			idExists = true
		}
	}

	if(idExists){
		Assignments[PK].Id =  r.FormValue("id")
		Assignments[PK].Title =  r.FormValue("title")
		Assignments[PK].Class =  r.FormValue("class")
		Assignments[PK].Description =  r.FormValue("desc")
		Assignments[PK].Points, _ =  strconv.Atoi(r.FormValue("points"))
		Assignments[PK].DueDate =  r.FormValue("duedate")
		Assignments[PK].TimeEstimate =  r.FormValue("timeestimate")
		w.WriteHeader(http.StatusCreated)
		return
	}
	w.WriteHeader(http.StatusNotFound)

}

func APIDeleteAssignment(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s DELETE end point", r.URL.Path)
	w.Header().Set("Content-Type", "application/txt")
	w.WriteHeader(http.StatusOK)
	
	response := make(map[string]string)

	if err := r.ParseForm(); err != nil {
		log.Print(w, "ParseForm() err: %v", err)
		return
	}

	PK, _ := strconv.Atoi(r.FormValue("PK"))

	for i, assignment := range Assignments {
		if assignment.PK == PK{
			Assignments = append(Assignments[:i], Assignments[i+1:]...)
			response["status"] = "Success"
			break
		}
	}
	
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return
	}
	w.Write(jsonResponse)
}

