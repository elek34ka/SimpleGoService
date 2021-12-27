package groupsHandler

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
	name := string(data)

	switch name {
	case "":
		data, err = json.Marshal(db.FindAllGroups())
		break
	default:
		data, err = json.Marshal(db.FindGroupByName(name))
		break
	}

	if err != nil {
		log.Println("Cant handle Get request for Group")
		return
	}

	ResponseHandler.SubmitResponse(w, data)
}

func Add(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)

	name := string(data)

	if err != nil {
		log.Println(w, "Cant handle Add request for Group")
		return
	}

	db.AddGroup(name)
	data, err = json.Marshal(db.FindGroupByName(name))
	ResponseHandler.SubmitResponse(w, data)
}
