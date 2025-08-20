package main

import (
	"log"
	"os"
	"time"
	"flag"
	"fmt"
)

//const filename = "README.md"

func main() {
	filename := flag.String("file", "./README.md", "target file")
	owner := flag.String("owner", "John Doe", "user")
	ignore := flag.String("ignore", "", "ignored repos")
	flag.Parse()

	now := time.Now().String()
	if err := os.WriteFile(*filename, []byte(*owner+"  \n"+*ignore+"  \n"+now+"\n"), 0644); err != nil {
		log.Fatal(err)
	}
}
