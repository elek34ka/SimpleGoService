package usersInGroupHandler

import (
	"SimpleGoService/db"
	"SimpleGoService/handler/ResponseHandler"
	"SimpleGoService/structs"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func Get(w http.ResponseWriter, r *http.Request) {
	data, _ := ioutil.ReadAll(r.Body)
	name := string(data)

	data, _ = json.Marshal(db.FindAllUsersInGroup(name))

	ResponseHandler.SubmitResponse(w, data)
}

func Add(w http.ResponseWriter, r *http.Request) {
	data, _ := ioutil.ReadAll(r.Body)

	var userInGroup structs.UserInGroup

	_ = json.Unmarshal(data, &userInGroup)

	db.AddUserInGroup(userInGroup)
}
