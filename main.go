package main

import (
	"log"
	"os"

	"github.com/amiraliio/tgbp-api/provider"
)

func main() {

	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}
	if err := os.Setenv("BOT_API_ROOT_DIR", currentDir); err != nil {
		log.Fatalln(err)
	}

	provider.HTTP()

	provider.GRPC()
}
