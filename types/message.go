package types

import (
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Message struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"ID,omitempty"`
	Session primitive.ObjectID `bson:"Session" json:"Session"`
	Sender  primitive.ObjectID `bson:"Sender" json:"Sender"`
	Content string             `bson:"Content" json:"Content"`
	Created time.Time          `bson:"Created" json:"Created"`
}
type CreateMessageParams struct {
	Session string    `json:"Session"`
	Sender  string    `json:"Sender"`
	Content string    `json:"Content"`
	Created time.Time `json:"Created"`
}

func (c *CreateMessageParams) FromParams() (*Message, error) {
	oidSession, err := primitive.ObjectIDFromHex(c.Session)
	if err != nil {
		return nil, fmt.Errorf("invalid session")
	}
	oidSender, err := primitive.ObjectIDFromHex(c.Sender)
	if err != nil {
		return nil, fmt.Errorf("invalid user")
	}
	if len(c.Content) < 1 {
		return nil, fmt.Errorf("empty message")
	}
	return &Message{
		Session: oidSession,
		Sender:  oidSender,
		Content: c.Content,
		Created: time.Now(),
	}, nil
}
