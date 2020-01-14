//Package routes ...
package routes

import (
	"github.com/amiraliio/tgbp-api/controller"
	"net/http"
	"github.com/gorilla/mux"
)

//Init ...
func Init(mux *mux.Router) {

	mux.HandleFunc("/user/{userID}/receiver/{receiverID}/direct-messages", controller.GetDirectMessageList).Methods(http.MethodGet)
}
