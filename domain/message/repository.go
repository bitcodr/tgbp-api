package message

//MessageRepository connect our bossiness logic to repository
type MessageRepository interface {
	DirectMessagesList(userID, receiverID, channelID int64) ([]*Message, error)
}
