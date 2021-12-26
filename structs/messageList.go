package structs

type Message struct {
	ID      int    `json:"id"`
	Sender  string `json:"sender"`
	Message string `json:"message"`
}

type MessageList []Message
