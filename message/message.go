package message

// Message is the object sent between the client and server to make requests.
type Message struct {
	Command    string   `json:"command"`
	Parameters []string `json:"parameters"`
}

// DropInsert is an insert request for a new drop.
type DropInsert struct {
	ObjectID uint32  `json:"object_id"`
	Rate     float32 `json:"rate"`
}
