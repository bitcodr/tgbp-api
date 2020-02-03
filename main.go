package main

import (
	"github.com/amiraliio/tgbp-api/repository/mysql"
	"github.com/amiraliio/tgbp-api/config"
	"github.com/amiraliio/tgbp-api/routes"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func main() {

	//get current directory
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	app := new(config.App)
	app.ProjectDir = currentDir

	app.Environment()

	repo := mysql.
	mux := mux.NewRouter()
	routes.Init(mux)
	http.ListenAndServe(":"+config.AppConfig.GetString("APP.PORT"), mux)
}



