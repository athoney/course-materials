package scrape



import (
	"regexp"
)


//==========================================================================\\
// || GLOBAL DATA STRUCTURES  ||

//ADVANCED: This is perhaps a terrible structure since multiple a filename is NOT guarenteed to be unique; consider an array of Locations? 
//CHALLENGE: Replace this Local Structure with a Key-Value DB like REDIS
type FileInfo struct {
	Id int `json:"id"`
	Filename string `json:"filename"`
	Location string `json:"location"`
}
var Files []FileInfo



var regexes = []*regexp.Regexp{
	regexp.MustCompile(`(?i)password`),
    regexp.MustCompile(`(?i).txt`),  
}

// END GLOBAL VARIABLES
//==========================================================================//

//==========================================================================\\
// || HELPER FUNCTIONS TO MANIPULATE THE REGULAR EXPRESSIONS ||

func resetRegEx(){
    regexes = []*regexp.Regexp{
        regexp.MustCompile(`(?i)password`),
        regexp.MustCompile(`(?i)kdb`),
        regexp.MustCompile(`(?i)login`),
    }
}

func clearRegEx(){
    regexes = nil
}

func addRegEx(regex string){
    regexes = append(regexes,regexp.MustCompile(regex))
}

//==========================================================================//




