package main

import (
	"log"
	"os"
	"time"
)

const filename = "README.md"

func main() {
	now := time.Now().String()
	if err := os.WriteFile(filename, []byte("\n"+now+"\n"), 0644); err != nil {
		log.Fatal(err)
	}
}
