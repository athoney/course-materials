package main

import (
	"fmt"
	"net/http"
	"math/rand" // Dr.Mike Adding for more fun examples
	"html/template"
	//"wyoassign/wyoassign"
)


type router struct {
}

func (r *router) ServeHTTP(w http.ResponseWriter, req *http.Request) { //(response, request)
	
	templates := template.Must(template.ParseFiles("home.html", "newAssignment.html", "assignments.html", "header.html","footer.html"))


	switch req.URL.Path {
	case "/":
		templates.ExecuteTemplate(w, "home", struct{ Name string}{"Alicia"})
	case "/newAssignment":
		templates.ExecuteTemplate(w, "newAssign", nil)
	case "/assignments":
		templates.ExecuteTemplate(w, "assign", nil)
	case "/index.html":
		fmt.Fprint(w, "<html><head><title>Real Website I Promise</title></head><body><H1>Welcome</H1>Nothing to see here</body></html>")
	case "/r":
		fmt.Fprint(w, "Hello World")
	case "/essay":
		bytes := make([]byte, 1000)
		for i := 0; i < 1000; i++ {
			bytes[i] = byte(rand.Intn(122-97))+97
		}
		fmt.Fprint(w, string(bytes))

	default:
		http.Error(w, "418 I'm a teapot", http.StatusTeapot)
	}
}

func main() {
	var r router
	http.ListenAndServe(":8000", &r)
}