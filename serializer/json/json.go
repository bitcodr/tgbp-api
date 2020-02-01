//Package json ...
package json

import (
	"encoding/json"
	"github.com/amiraliio/tgbp-api/domain/message"
	"github.com/pkg/errors"
)

type Message struct{}

func (m *Message) Encode(input *message.Message) ([]byte, error) {
	rawMessage, err := json.Marshal(input)
	if err != nil {
		return nil, errors.Wrap(err, "serializer.Message.Encode")
	}
	return rawMessage, nil
}

func (m *Message) Decode(input []byte) (*message.Message, error) {
	message := new(message.Message)
	if err := json.Unmarshal(input, message); err != nil {
		return nil, errors.Wrap(err, "serializer.Message.Decode")
	}
	return message, nil
}
