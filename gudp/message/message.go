package message

type Message struct {
	msg    []byte
	client string
}

func NewMessage(msg []byte, client string) *Message {
	return &Message{
		msg:    msg,
		client: client,
	}
}

func (m *Message) GetMessage() []byte {

	return m.msg
}

func (m *Message) GetClient() string {
	return m.client
}
