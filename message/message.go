package message

// Message is the object sent between the client and server to make requests.
type Message struct {
	Command    string   `json:"command"`
	Parameters []string `json:"parameters"`
}

// Result is the result of an action comitted by this application.
type Result struct {
	Success bool   `json:"success"`
	Reason  string `json:"reason"`
}
