package ResponseHandler

import (
	"log"
	"net/http"
)

func SubmitResponse(w http.ResponseWriter, data []byte) {
	log.Println("Return data")
	w.Header().Add("Content-Type", "application/json")
	w.Write(data)
}
