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
	fmt.Printf("=========== %v ===========", filename)
	now := time.Now().String()
	if err := os.WriteFile(*filename, []byte("\n"+now+"\n"), 0644); err != nil {
		log.Fatal(err)
	}
}
