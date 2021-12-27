package usersHandler

import (
	"SimpleGoService/db"
	"SimpleGoService/handler/ResponseHandler"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func Get(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	username := string(data)

	switch username {
	case "":
		data, err = json.Marshal(db.FindAllUsers())
		break
	default:
		data, err = json.Marshal(db.FindUserByName(username))
		break
	}

	if err != nil {
		log.Println("Cant handle Get request for Users")
		return
	}

	ResponseHandler.SubmitResponse(w, data)
}

func Add(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)

	username := string(data)

	if err != nil {
		log.Println(w, "Cant handle Add request for Users")
		return
	}

	db.AddUser(username)
	data, err = json.Marshal(db.FindUserByName(username))
	ResponseHandler.SubmitResponse(w, data)
}
