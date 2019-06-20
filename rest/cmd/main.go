package main

import (
	app "github.com/alemjc/gophercises/cyoa/rest/api"
	"log"
	"os"
	"strconv"
)

func main() {
	port, err := strconv.Atoi(os.Getenv("PORT"))

	if err != nil {
		log.Fatal(err)
		return
	}

	log.Fatal(app.Run(port))

}
