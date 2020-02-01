package message

//MessageRepository connect our bussiness logic to repository
type MessageRepository interface {
	DirectMessagesList(userID, receiverID, channelID int64) ([]*Message, error)
}
