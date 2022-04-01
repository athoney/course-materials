package scrape

// scrapeapi.go HAS TEN TODOS - TODO_5-TODO_14 and an OPTIONAL "ADVANCED" ASK

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"github.com/gorilla/mux"
)

//==========================================================================\\

// Helper function walk function, modfied from Chap 7 BHG to enable passing in of
// additional parameter http responsewriter; also appends items to global Files and
// if responsewriter is passed, outputs to http
var LOG_LEVEL int
var index = 0
var rootdir = `/home/cabox/`

func walkFn(w http.ResponseWriter) filepath.WalkFunc {
    return func(path string, f os.FileInfo, err error) error {
        w.Header().Set("Content-Type", "application/json")

        for _, r := range regexes {
            if r.MatchString(path) {
                var tfile FileInfo
                dir, filename := filepath.Split(path)
                tfile.Filename = string(filename)
                tfile.Location = string(dir)
				tfile.Id = index

				var exists = false
				for _, file := range Files{
					if file.Filename == tfile.Filename && file.Location == tfile.Location{
						exists = true
					}
				}
                if (!exists) {
					Files = append(Files, tfile)
					index++

				}

                if w != nil && len(Files)>0 {
                    w.Write([]byte(`"`+ strconv.Itoa(tfile.Id) +`":  `))
                    json.NewEncoder(w).Encode(tfile)
                    w.Write([]byte(`,`))

                } 
                if LOG_LEVEL == 2 {
					log.Printf("[+] HIT: %s\n", path)
				}

            }

        }
        return nil
    }

}


func walkFn2(w http.ResponseWriter, query string) filepath.WalkFunc {
    return func(path string, f os.FileInfo, err error) error {
		w.Header().Set("Content-Type", "application/json")

		r := regexp.MustCompile(`(?i)`+query)

		if r.MatchString(path) {
			var tfile FileInfo
			dir, filename := filepath.Split(path)
			tfile.Filename = string(filename)
			tfile.Location = string(dir)
			tfile.Id = index

			var exists = false
			for _, file := range Files{
				if file.Filename == tfile.Filename && file.Location == tfile.Location{
					exists = true
				}
			}
			if (!exists) {
				Files = append(Files, tfile)
				index++

			}

			if w != nil && len(Files)>0 {
				w.Write([]byte(`"`+ strconv.Itoa(tfile.Id) +`":  `))
				json.NewEncoder(w).Encode(tfile)
				w.Write([]byte(`,`))

			} 
			if LOG_LEVEL == 2 {
				log.Printf("[+] HIT: %s\n", path)
			}

		}

	return nil
		
    }
}

//==========================================================================\\

func APISTATUS(w http.ResponseWriter, r *http.Request) {

	if LOG_LEVEL >= 1 {
		log.Printf("Entering %s end point", r.URL.Path)
	}
	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(`{ "status" : "API is up and running ",`))
    var regexstrings []string
    
    for _, regex := range regexes{
        regexstrings = append(regexstrings, regex.String())
    }

    w.Write([]byte(` "regexs" :`))
    json.NewEncoder(w).Encode(regexstrings)
    w.Write([]byte(`}`))
	log.Println(regexes)

}


func MainPage(w http.ResponseWriter, r *http.Request) {
	if LOG_LEVEL >= 1 {
		log.Printf("Entering %s end point", r.URL.Path)
	}
    w.Header().Set("Content-Type", "text/html")

	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, "<html><body><H1>Welcome to my awesome File page</H1><ol><li><b>/</b> - Current location</li><li><b>/api-status</b> - see info on current regexs</li><li><b>/indexer</b> - search for files</li><li><b>/search</b> - find a file</li><li><b>/addsearch/{regex}</b> - add regex to search on</li><li><b>/clear</b> - clear regexs</li><li><b>/reset</b> - reset to default regex list</li></body>")
}

//Fix"
func FindFile(w http.ResponseWriter, r *http.Request) {
	if LOG_LEVEL >= 1 {
		log.Printf("Entering %s end point", r.URL.Path)
	}
    q, ok := r.URL.Query()["q"]

    w.WriteHeader(http.StatusOK)
    if ok && len(q[0]) > 0 {
		if LOG_LEVEL >= 1 {
        	log.Printf("Entering search with query=%s",q[0])
		}

        // ADVANCED: Create a function in scrape.go that returns a list of file locations; call and use the result here
        // e.g., func finder(query string) []string { ... }
		var found = false
        for _, File := range Files {
		    if File.Filename == q[0] {
                json.NewEncoder(w).Encode(File.Location)
                found = true
		    }
        }
		if !found {
			fmt.Fprintf(w, "<html><h1>File not found</h1>")
		}
    } else {
        // didn't pass in a search term, show all that you've found
        w.Write([]byte(`"files":`))    
        json.NewEncoder(w).Encode(Files)
    }
}

func IndexFiles(w http.ResponseWriter, r *http.Request) {
	if LOG_LEVEL >= 1 {
    	log.Printf("Entering %s end point", r.URL.Path)
	}
    w.Header().Set("Content-Type", "application/json")

    location, locOK := r.URL.Query()["location"]

    if locOK && len(location[0]) > 0 {
        w.WriteHeader(http.StatusOK)

    } else {
        w.WriteHeader(http.StatusFailedDependency)
        w.Write([]byte(`{ "parameters" : {"required": "location",`))    
        w.Write([]byte(`"optional": "regex"},`))    
        w.Write([]byte(`"examples" : { "required": "/indexer?location=/xyz",`))
        w.Write([]byte(`"optional": "/indexer?location=/xyz&regex=(i?).md"}}`))
        return 
    }

    //wrapper to make "nice json"
    w.Write([]byte(`{ `))
    
    regex, regexOK := r.URL.Query()["regex"]
    
    if regexOK {
		filepath.Walk(rootdir+location[0], walkFn2(w, regex[0]))
	} else {
		filepath.Walk(rootdir+location[0], walkFn(w))
	}

    //wrapper to make "nice json"
    w.Write([]byte(` "status": "completed"} `))

}


func ResetArrs(w http.ResponseWriter, r *http.Request) {
	if LOG_LEVEL >= 1 {
    	log.Printf("Entering %s end point", r.URL.Path)
	}
    w.Header().Set("Content-Type", "application/json")
	resetRegEx()
	index = 0
	Files = nil
}

func Clear(w http.ResponseWriter, r *http.Request) {
	if LOG_LEVEL >= 1 {
    	log.Printf("Entering %s end point", r.URL.Path)
	}
    w.Header().Set("Content-Type", "application/json")
	clearRegEx()
}

func AddRegex(w http.ResponseWriter, r *http.Request) {
	if LOG_LEVEL >= 1 {
    	log.Printf("Entering %s end point", r.URL.Path)
	}
    w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	addRegEx(`(?i)`+params["regex"])
}
