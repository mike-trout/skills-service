// main.go

package main

import (
	"log"
	"os"
)

func main() {
	a := App{}
	log.Println("Initialising app")
	a.Initialise()
	addr := os.Getenv("APP_ADDR")
	log.Printf("Running app on %s\n", addr)
	a.Run(addr)
}
