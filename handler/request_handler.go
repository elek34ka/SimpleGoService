package handler

import (
	"SimpleGoService/storage"
	"SimpleGoService/structs"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func Get(w http.ResponseWriter, r *http.Request) {
	var data, err = json.Marshal(storage.Get())

	if err != nil {
		HandleError(w, "Cant handle PostRequest")
		return
	}

	log.Println("Return data")
	w.Header().Add("Content-Type", "application/json")
	w.Write(data)
}

func Post(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)

	var mes structs.Message

	err = json.Unmarshal(data, &mes)

	if err != nil {
		HandleError(w, "Cant handle PostRequest")
		return
	}

	id := storage.Add(mes)

	log.Println("Added your message ", mes, " with id = ", id)

}

func Delete(w http.ResponseWriter, r *http.Request) {
	var data, err = ioutil.ReadAll(r.Body)

	type mesId struct {
		ID int `json:"id"`
	}

	var id mesId
	err = json.Unmarshal(data, &id)

	if err != nil {
		HandleError(w, "Cant handle DeleteRequest")
		return
	}

	var mes = storage.Delete(id.ID)

	log.Println("Deleted message ", mes)
}

func Default(w http.ResponseWriter, r *http.Request) {
	var data = "default request"
	w.Header().Add("Content-Type", "application/json")
	w.Write([]byte(data))
}
