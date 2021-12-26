package handler

import (
	"log"
	"net/http"
)

func HandleError(w http.ResponseWriter, err string) {
	w.Header().Add("Content-Type", "application/json")
	w.Write([]byte(err))
}

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	log.Println("Request received")
	switch r.Method {
	case http.MethodGet:
		Get(w, r)
		break
	case http.MethodPost:
		Post(w, r)
		break
	case http.MethodDelete:
		Delete(w, r)
		break
	default:
		Default(w, r)
		break
	}
}
