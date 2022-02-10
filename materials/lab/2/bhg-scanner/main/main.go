package main

import (
	"fmt"
	"bhg-scanner/scanner"
)

func main(){
	var start int
	var end int

	fmt.Printf("Enter the port to begin scanning at: ")
	fmt.Scanln(&start)
	fmt.Printf("Enter the port to stop scanning at: ")
	fmt.Scanln(&end)

	scanner.PortScanner(start, end)
}