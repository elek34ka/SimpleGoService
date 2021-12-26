package storage

import "my_first_service/structs"

var store structs.MessageList
var curMaxId = 1

func Get() structs.MessageList {
	return store
}

func Add(m structs.Message) int {
	m.ID = curMaxId
	curMaxId++
	store = append(store, m)
	return m.ID
}

func Delete(id int) structs.Message {
	ind := -1
	var message structs.Message
	for i, mes := range store {
		if mes.ID == id {
			ind = i
			message = mes
		}
	}

	if ind != -1 {
		store = append(store[:ind], store[ind+1:]...)
	}
	
	return message
}
