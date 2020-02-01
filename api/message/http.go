//Package message ...
package message

import (
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/amiraliio/tgbp-api/serializer/json"
	"github.com/amiraliio/tgbp-api/domain/message"
	"net/http"
	"html/template"
	"log"
	"strconv"
)

type MessageHandler interface {
	GetDirectMessages(res http.ResponseWriter, req *http.Request)
}

type Handler struct{
    messageService message.MessageService
}


func NewHandler(messageService message.MessageService) MessageHandler{
	return &Handler{messageService: messageService}
}

func setupResponse(res http.ResponseWriter, contentType string, body []byte, statusCode int){
	res.Header().Set("Content-Type",contentType)
	res.WriteHeader(statusCode)
	_, err := res.Write(body)
	if err !=nil{
		log.Println(err)
	}
}


func (h *handler) serializer(contentType string) message.MessageSerializer{
	switch contentType{
	case "application/json":
		return &json.Message{}
	}
}

func (h *Handler) GetDirectMessages(res http.ResponseWriter, req *http.Request){
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
	if err !=nil{
		if errors.Cause(err) == message.ErrMessageNotFound{
			http.Error(res, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		http.Error(res, http.StatusText(http.StatusInternalServerError),http.StatusInternalServerError)
		return
	}
	tpl := template.Must(template.ParseFiles("resource/template/chat.html"))
	if err := tpl.Execute(res, allDM); err != nil {
		log.Println(err)
		return
	}
}
