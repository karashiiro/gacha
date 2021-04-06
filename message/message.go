package message

// Message is the object sent between the client and server to make requests.
type Message struct {
	Command   string `json:"command"`
	Arguments string `json:"arguments"`
}
