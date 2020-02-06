package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/amiraliio/tgbp-api/config"
	"github.com/amiraliio/tgbp-api/domain/message"
	"github.com/amiraliio/tgbp-api/handler/web"
	"github.com/amiraliio/tgbp-api/repository/arango"
	"github.com/amiraliio/tgbp-api/repository/mysql"
	"github.com/gorilla/mux"
)

func main() {

	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}
	if err := os.Setenv("BOT_API_ROOT_DIR", currentDir); err != nil {
		log.Fatalln(err)
	}
	app := new(config.App)
	app.Environment()
	db := app.DB()

	messageRepo := chooseMessageRepo("mysql", db, app)

	messageService := message.NewMessageService(messageRepo)

	messageHandler := web.NewWebMessageHandler(messageService)

	router := mux.NewRouter()

	router.HandleFunc("/user/{userID}/receiver/{receiverID}/channel/{channelID}/direct-messages", messageHandler.GetDirectMessages).Methods(http.MethodGet)

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

func chooseMessageRepo(connection string, db *sql.DB, appConfig *config.App) message.MessageRepository {
	switch connection {
	case "mysql":
		return mysql.NewMysqlMessageRepository(db, appConfig)
	case "arango":
		return arango.NewArangoMessageRepository(appConfig)
	default:
		return nil
	}
}
