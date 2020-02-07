//Package rest ...
package rest

import (
	"net/http"

	domain "github.com/amiraliio/tgbp-api/domain/message"
	"github.com/amiraliio/tgbp-api/serializer/json"
	"github.com/amiraliio/tgbp-api/serializer/msgpack"
)

type MessageHandler interface {
	GetDirectMessages(res http.ResponseWriter, req *http.Request)
}

type messageHandler struct {
	messageService domain.MessageService
}

func NewRestMessageHandler(messageService domain.MessageService) MessageHandler {
	return &messageHandler{
		messageService,
	}
}

func (h *messageHandler) serializer(contentType string) domain.MessageSerializer {
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
