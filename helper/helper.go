//Package helper ...
package helper

import (
	"log"
	"net/http"
)

func SetupResponse(res http.ResponseWriter, contentType string, body []byte, statusCode int) {
	res.Header().Set("Content-Type", contentType)
	res.WriteHeader(statusCode)
	if _, err := res.Write(body); err != nil {
		log.Println(err)
	}
}
