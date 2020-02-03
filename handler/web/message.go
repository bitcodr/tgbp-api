//Package web ...
package web

import (
	"github.com/amiraliio/tgbp-api/domain/message"
	"github.com/amiraliio/tgbp-api/handler"
	"github.com/amiraliio/tgbp-api/serializer/json"
	"github.com/amiraliio/tgbp-api/serializer/msgpack"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type messageHandler struct {
	messageService message.MessageService
}

func NewWebMessageHandler(messageService message.MessageService) handler.MessageHandler {
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
	params := mux.Vars(req)
	if params == nil {
		log.Println("params are nil")
		return
	}
	userID, err := strconv.ParseInt(params["userID"], 10, 0)
	if err != nil {
		log.Println(err)
		return
	}
	receiverID, err := strconv.ParseInt(params["receiverID"], 10, 0)
	if err != nil {
		log.Println(err)
		return
	}
	channelID, err := strconv.ParseInt(params["channelID"], 10, 0)
	if err != nil {
		log.Println(err)
		return
	}
	messages, err := h.messageService.DirectMessagesList(userID, receiverID, channelID)
	if err != nil {
		if errors.Cause(err) == message.ErrMessageNotFound {
			http.Error(res, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		http.Error(res, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	tpl := template.Must(template.ParseFiles("resource/template/chat.html"))
	if err := tpl.Execute(res, messages); err != nil {
		log.Println(err)
		return
	}
}
