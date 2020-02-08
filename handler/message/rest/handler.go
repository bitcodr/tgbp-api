//Package rest ...
package rest

import (
	"net/http"

	"github.com/amiraliio/tgbp-api/domain/service"
	"github.com/amiraliio/tgbp-api/serializer/json"
	"github.com/amiraliio/tgbp-api/serializer/msgpack"
)

type MessageHandler interface {
	GetDirectMessages(res http.ResponseWriter, req *http.Request)
}

type messageHandler struct {
	messageService service.MessageService
}

func NewRestMessageHandler(messageService service.MessageService) MessageHandler {
	return &messageHandler{
		messageService,
	}
}

func (h *messageHandler) serializer(contentType string) service.MessageSerializer {
	switch contentType {
	case "application/json":
		return &json.Message{}
	case "application/x-msgpack":
		return &msgpack.Message{}
	default:
		return &json.Message{}
	}
}

func (h *messageHandler) GetDirectMessages(res http.ResponseWriter, req *http.Request) {
	//TODO
}
