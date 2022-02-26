# Lab 3b - Shodan API
## To run this file: 
1. navigate to /course-materials/materials/lab/3/main
2. type: `go build main.go`
3. type: `SHODAN_API_KEY={your API key} ./main {query}`
4. follow command prompts
## Example Queries:
- `city:Laramie`
- `org:Apple`
- `city:Laramie org:"University of Wyoming"`
## Purpose:
This program utilizies Shodan's APIs to search the Shodan website with a query from the command line.
## Modifications:
1. I modified main.go so that the user can page through their query results by continuing to type "Y" in the command line. I also added the ISP of each result to the output of the program.
