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
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	err := godotenv.Load() //Load env.file
	if err != nil {
		log.Fatalln(errors.Wrap(err, "failed to load .env file at godotenv.Load()"))
	}

	e := infrastructure.NewRouter()

	//infrastructure.E.Run(os.Getenev("API_SERVER_PORT"))
	log.Println("server start...")
	e.Logger.Fatal(e.Start(fmt.Sprintf("127.0.0.1:%s", os.Getenv("API_SERVER_PORT"))))
}
