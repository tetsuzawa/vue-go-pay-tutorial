package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"log"
	"os"

	"github.com/tetsuzawa/vue-go-pay-tutorial/backend-api/infrastructure"
)

func main() {
	err := godotenv.Load() //Load env.file
	if err != nil {
		log.Fatalln(errors.Wrap(err, "failed to load .env file at godotenv.Load()"))
	}

	//infrastructure.E.Run(os.Getenev("API_SERVER_PORT"))
	infrastructure.E.Logger.Fatal(infrastructure.E.Start(fmt.Sprintf(":%s", os.Getenv("API_SERVER_PORT"))))
}
