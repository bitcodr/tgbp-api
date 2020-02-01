package message

type MessageSerializer interface {
	Encode(input *Message) ([]byte, error)
	Decode(input []byte) (*Message, error)
}
