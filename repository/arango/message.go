package arango

import "github.com/amiraliio/tgbp-api/domain/message"

type messageRepo struct{}


func NewArangoMessageRepository() message.MessageRepository{
	return &messageRepo{}
}


func (a *messageRepo) DirectMessagesList(userID, receiverID, channelID int64) ([]*message.Message, error) {
	panic("implement me")
}