package main

import (
	"fmt"
	"hscan/hscan"
	"log"
	"os"
)

func main() {

	if len(os.Args) != 2 {
		log.Fatalln("Usage: main <wordlist file name>")
	}

	fileName := os.Args[1]


	//To test this with other password files youre going to have to hash
	var md5hash = "77f62e3524cd583d698d51fa24fdff4f"
	var sha256hash = "95a5e1547df73abdd4781b6c9e55f3377c15d08884b11738c2727dbd887d4ced"
	var sha256hash2 = "e7cf3ef4f17c3999a94f2c6f612e8a888e5b1026878e4e19398b23bd38ec221a"


	var drmike1 = "90f2c9c53f66540e67349e0ab83d8cd0" //p@ssword
	var drmike2 = "1c8bfe8f801d79745c4631d09fff36c82aa37fc4cce4fc946683d7b336b63032" //letmeins

	var file = "/home/cabox/workspace/course-materials/materials/lab/7/main/"+fileName
	
	
	hscan.GuessSingle(md5hash, file)
	hscan.GuessSingle(sha256hash, file)
	hscan.GuessSingle(sha256hash2, file)

	fmt.Printf("Dr.Mike 1: %s\n", hscan.GuessSingle(drmike1, file))
	fmt.Printf("Dr.Mike 2: %s\n", hscan.GuessSingle(drmike2, file))
	
	hscan.GenHashMaps(file)
	
	hscan.GetSHA(sha256hash)
	hscan.GetMD5(sha256hash)
}
