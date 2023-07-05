package chat

import (
	"github.com/google/uuid"
)

type Msg struct {
	Id        uuid.UUID `bson:"id"`
	Message   string    `bson:"message"`
	recipient *Client
}

func NewMessage(client *Client, text string) *Msg {
	uuid, err := uuid.NewRandom()
	if err != nil {
		return nil
	}

	return &Msg{
		Id:        uuid,
		Message:   text,
		recipient: client,
	}
}

func (m *Msg) GetMessage() []byte {
	return []byte(m.Message)
}

func (m *Msg) SentBy(reciever *Client) bool {
	return m.recipient == reciever
}
