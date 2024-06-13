package main

import (
	"github.com/Aries-Financial-inc/golang-dev-logic-challenge-sulavneupane/api"
	_ "github.com/joho/godotenv/autoload"
	"log"
)

func main() {
	apiImplementation := api.NewBuilder().
		WithRepositoryProvider(api.GetRepositories()).
		Finalize()

	if err := apiImplementation.ListenAndServe(); err != nil {
		log.Fatalln("ListenAndServe Failed with Fatal Error: ", err)
	}
}
