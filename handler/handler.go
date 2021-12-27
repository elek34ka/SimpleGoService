package handler

import (
	"SimpleGoService/handler/groupsHandler"
	"SimpleGoService/handler/usersHandler"
	"log"
	"net/http"
)

func HandleError(w http.ResponseWriter, err string) {
	w.Header().Add("Content-Type", "application/json")
	w.Write([]byte(err))
}

func CommonHandler() {
	http.HandleFunc("/users", UsersHandler)
	http.HandleFunc("/groups", GroupsHandler)
	http.HandleFunc("/", HandleRequest)
}

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Users Request received")
	switch r.Method {
	case http.MethodGet:
		usersHandler.Get(w, r)
		break
	case http.MethodPost:
		usersHandler.Add(w, r)
		break
	default:
		HandleRequest(w, r)
		break
	}
}

func GroupsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Groups Request received")
	switch r.Method {
	case http.MethodGet:
		groupsHandler.Get(w, r)
		break
	case http.MethodPost:
		groupsHandler.Add(w, r)
		break
	default:
		HandleRequest(w, r)
		break
	}
}

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	log.Println("Request received")
}
