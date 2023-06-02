package message

type Message string

func NewMessage() Message {
	return "Hi there!"
}

func NewCustomMessage(msg string) Message {
	return Message(msg)
}
