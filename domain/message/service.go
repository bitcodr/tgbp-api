package message

//MessageService 
type MessageService interface {
	DirectMessagesList(userID, receiverID, channelID int64) ([]*Message, error)
}
