//Package provider ...
package provider

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/amiraliio/tgbp-api/config"
	message "github.com/amiraliio/tgbp-api/handler/message"
	"github.com/gorilla/mux"
)

func HTTP() {
	app := new(config.App)
	app.Init()

	router := mux.NewRouter()

	//Implements routes in here
	message.Routes(app, router)

	errs := make(chan error, 2)

	go func() {
		fmt.Println("Listening on port " + config.AppConfig.GetString("APP.PORT"))
		errs <- http.ListenAndServe(":"+config.AppConfig.GetString("APP.PORT"), router)
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	fmt.Printf("terminated %s", <-errs)

}

func GRPC() {
	//TODO GRPC handlers
}
