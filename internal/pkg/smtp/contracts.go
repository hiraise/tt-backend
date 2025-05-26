package smtp

type Sender interface {
	Send(msg Message, eventId string) error
}

type Message struct {
	Subject    string
	Recipients []string
	Text       string
}
