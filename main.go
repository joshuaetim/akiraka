package main

import (
	"log"

	"github.com/joho/godotenv"
	route "github.com/joshuaetim/quiz/route"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(route.RunAPI(":5000"))
}
