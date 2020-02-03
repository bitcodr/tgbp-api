//Package handler ...
package handler

import "net/http"

type MessageHandler interface {
	GetDirectMessages(res http.ResponseWriter, req *http.Request)
}