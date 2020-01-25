//Package routes ...
package routes

import (
	"github.com/amiraliio/tgbp-api/controller"
	"github.com/gorilla/mux"
	"net/http"
)

//Init ...
func Init(mux *mux.Router) {

	mux.HandleFunc("/user/{userID}/receiver/{receiverID}/channel/{channelID}/direct-messages", controller.GetDirectMessageList).Methods(http.MethodGet)
}
