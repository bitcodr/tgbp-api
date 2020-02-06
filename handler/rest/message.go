//Package rest ...
package rest

import (
	"github.com/amiraliio/tgbp-api/domain/message"
	"github.com/amiraliio/tgbp-api/handler"
	"github.com/amiraliio/tgbp-api/serializer/json"
	"github.com/amiraliio/tgbp-api/serializer/msgpack"
	"net/http"
)

type messageHandler struct {
	messageService message.MessageService
}

func NewRestMessageHandler(messageService message.MessageService) handler.MessageHandler {
	return &messageHandler{
		messageService,
	}
}

func (h *messageHandler) serializer(contentType string) message.MessageSerializer {
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
