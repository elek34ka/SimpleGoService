package groupsHandler

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

	SubmitResponse(w, data)
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
	SubmitResponse(w, data)
}
