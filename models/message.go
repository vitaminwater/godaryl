package models

type Message struct {
	Id   string    `db:"id" json:"id"`
	Text string    `json:"text,omitempty"`
	At   time.Time `db:"at" json:"at"`
}

type LinkMessage struct {
	Message
	Link string `db:"link" json:"link,omitempty"`
}

type NoteMessage struct {
	Message
}

type TodoMessage struct {
	Message
	Done bool   `db:"done" json:"done,omitempty"`
	Todo string `db:"todo" json:"todo,omitempty"`
}
