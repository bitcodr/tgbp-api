//Package msgpack ...
package msgpack

import (
	"github.com/amiraliio/tgbp-api/domain/message"
	"github.com/pkg/errors"
	"github.com/vmihailenco/msgpack/v4"
)

type Message struct{}

func (m *Message) Encode(input *message.Message) ([]byte, error) {
	rawMessage, err := msgpack.Marshal(input)
	if err != nil {
		return nil, errors.Wrap(err, "serializer.Message.Encode")
	}
	return rawMessage, nil
}

func (m *Message) Decode(input []byte) (*message.Message, error) {
	messageModel := new(message.Message)
	if err := msgpack.Unmarshal(input, messageModel); err != nil {
		return nil, errors.Wrap(err, "serializer.Message.Decode")
	}
	return messageModel, nil
}
