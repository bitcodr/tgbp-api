package main

import (
	"github.com/amiraliio/tgbp-api/config"
	"github.com/amiraliio/tgbp-api/routes"
	"net/http"
)

func main() {

	app := new(config.App)
	app.SetAppConfig()
	app.Environment()

	mux := http.NewServeMux()
	routes.Init(mux)
	http.ListenAndServe(":"+config.AppConfig.GetString("APP.PORT"), mux)
}
