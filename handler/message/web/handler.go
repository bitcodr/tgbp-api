//Package web ...
package web

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/amiraliio/tgbp-api/config"
	"github.com/amiraliio/tgbp-api/domain/service"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

type MessageHandler interface {
	GetDirectMessages(res http.ResponseWriter, req *http.Request)
}

type messageHandler struct {
	messageService service.MessageService
}

func NewWebMessageHandler(messageService service.MessageService) MessageHandler {
	return &messageHandler{
		messageService,
	}
}

func (h *messageHandler) GetDirectMessages(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	if params == nil {
		log.Println(config.LangConfig.GetString("MESSAGES.PARAM_EMPTY"))
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
		if errors.Cause(err) == service.ErrMessageNotFound {
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
