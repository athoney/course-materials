// bhg-scanner/scanner.go modified from Black Hat Go > CH2 > tcp-scanner-final > main.go
// Code : https://github.com/blackhat-go/bhg/blob/c27347f6f9019c8911547d6fc912aa1171e6c362/ch-2/tcp-scanner-final/main.go
// License: {$RepoRoot}/materials/BHG-LICENSE
// Useage: navigate to {reporoot}/materials/lab/2/bhg-scanner/main and enter "go build main.go" and run "./main"

package scanner

import (
	"fmt"
	"net"
	"sort"
	"time"
)

var closedports []int
var openports []int


func worker(ports, results chan int) {
	for p := range ports {
		address := fmt.Sprintf("scanme.nmap.org:%d", p)    
		conn, err := net.DialTimeout("tcp", address, 2 * time.Second)
		if err != nil { 
			results <- (-1 * p)
			continue
		}
		conn.Close()
		results <- p
	}
}

// for Part 5 - consider
// easy: taking in a variable for the ports to scan (int? slice? ); a target address (string?)?
// med: easy + return  complex data structure(s?) (maps or slices) containing the ports.
// hard: restructuring code - consider modification to class/object 
// No matter what you do, modify scanner_test.go to align; note the single test currently fails
func PortScanner() (int, int) {  

	var start int
	var end int

	fmt.Println("Enter the port to begin scanning at: ")
	fmt.Scanln(&start)
	fmt.Println("Enter the port to stop scanning at: ")
	fmt.Scanln(&end)


	ports := make(chan int, 450)
	results := make(chan int)

	for i := 0; i < cap(ports); i++ {
		go worker(ports, results)
	}

	go func() {
		for i := start; i <= end; i++ {
			ports <- i
		}
	}()

	for i := 0; i < end; i++ {
		port := <-results
		if port >= 0 {
			openports = append(openports, port)
		} else {
			closedports = append(closedports, (port * -1))
		}
	}

	close(ports)
	close(results)
	sort.Ints(openports)
	sort.Ints(closedports)

	// for _, port := range openports {
	// 	fmt.Printf("%d, OPEN\n", port)
	// }

	for _, port := range closedports {
		fmt.Printf("%d, closed\n", port)
	}

	return len(openports), len(closedports)
}