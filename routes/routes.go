//Package routes ...
package routes

import (
	"github.com/amiraliio/tgbp-api/controller"
	"net/http"
)

//Init ...
func Init(mux *http.ServeMux) {

    mux.HandleFunc("/user/:userID/receiver/:receiverID/direct-messages",  controller.GetDirectMessageList)
}

//https://freshman.tech/web-development-with-go/