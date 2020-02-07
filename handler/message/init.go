//Package message ...
package message

import (
	"net/http"

	"github.com/amiraliio/tgbp-api/config"
	domain "github.com/amiraliio/tgbp-api/domain/message"
	web "github.com/amiraliio/tgbp-api/handler/message/web"
	"github.com/amiraliio/tgbp-api/repository/arango"
	"github.com/amiraliio/tgbp-api/repository/mysql"
	"github.com/gorilla/mux"
)

//Route names
const (
	WEB_DIRECT_MESSAGES_HISTORY = "WEB_DIRECT_MESSAGES_HISTORY"
)

func chooseMessageRepo(connection string, app *config.App) domain.MessageRepository {
	switch connection {
	case "mysql":
		return mysql.NewMysqlMessageRepository(app)
	case "arango":
		return arango.NewArangoMessageRepository(app)
	default:
		return nil
	}
}

func HTTP(app *config.App, router *mux.Router) {

	messageRepo := chooseMessageRepo("mysql", app)

	messageService := domain.NewMessageService(messageRepo)

	messageWebHandler := web.NewWebMessageHandler(messageService)

	router.HandleFunc("/user/{userID}/receiver/{receiverID}/channel/{channelID}/direct-messages", messageWebHandler.GetDirectMessages).Methods(http.MethodGet).Name(WEB_DIRECT_MESSAGES_HISTORY)
}

func GRPC(app *config.App) {
	//implement grpc handler route here
}
