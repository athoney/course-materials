# Lab 4 - Shodan API
## To run this file: 
1. navigate to /course-materials/materials/lab/3/main
2. type: `go build main.go`
3. type: `SHODAN_API_KEY={your API key} ./main {query}`
4. follow command prompts
## Purpose:
Main utilizies Shodan's APIs to search the Shodan website with a query from the command line.
## Modifications:
1. I created myip.go which uses the utility method /tools/myip to return your *not so exciting* IP address as seen by the internet. This is a simple modification, but I enjoyed the practice.
2. I modified host.go to include a page number variable in the query
3. I modified main.go so that the user can page through their query results by continuing to type "Y" in the command line. I also added the ISP of each result to the output of the program.
