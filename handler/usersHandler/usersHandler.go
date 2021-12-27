package usersHandler

import (
	"SimpleGoService/db"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func SubmitResponse(w http.ResponseWriter, data []byte) {
	log.Println("Return data")
	w.Header().Add("Content-Type", "application/json")
	w.Write(data)
}

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

	SubmitResponse(w, data)
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
	SubmitResponse(w, data)
}
