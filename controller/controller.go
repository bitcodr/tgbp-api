//Package controller ...
package controller

import (
	"github.com/amiraliio/tgbp-api/repository"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func getRepo() repository.Repo {
	return new(repository.RepoService)
}

//TODO check the userID and receiverId exist or not in the map

func GetDirectMessageList(res http.ResponseWriter, req *http.Request) {
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
	allDM, err := getRepo().GetAllDM(userID, receiverID)
	if err != nil {
		log.Println(err)
		return
	}
	tpl := template.Must(template.ParseFiles("template/chat.html"))
	if err := tpl.Execute(res, allDM); err != nil {
		log.Println(err)
		return
	}
}
