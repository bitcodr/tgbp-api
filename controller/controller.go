//Package controller ...
package controller

import (
	"github.com/amiraliio/tgbp-api/repository"
	"net/http"
)

func getRepo() repository.Repo {
	return new(repository.RepoService)
}

type Service struct{}

func (service *Service) GetDirectMessageList(res http.ResponseWriter, req *http.Request) {

}
