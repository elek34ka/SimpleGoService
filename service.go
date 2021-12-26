package main

import (
	"SimpleGoService/db"
	"SimpleGoService/structs"
	"log"
)

const PORT = 1488

func createMessage(s1 string, s2 string) structs.Message {
	return structs.Message{
		Sender:  s1,
		Message: s2,
	}
}

func main() {
	log.Println("Started program")

	db.Run()
	//storage.Add(createMessage("ruslan", "123"))
	//storage.Add(createMessage("maxim", "12"))
	//storage.Add(createMessage("artwem", "1"))
	//storage.Add(createMessage("zloyTapok", "89"))
	//
	//http.HandleFunc("/", handler.HandleRequest)
	//
	//var err = http.ListenAndServe(":"+strconv.Itoa(PORT), nil)
	//
	//if err != nil {
	//	log.Panicln("Server failed starting. Error: %s", err)
	//}
}
