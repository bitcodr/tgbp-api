//Package message ...
package message

import "errors"

//repo <-via aport<- service -> serializer -> http
//validate of data should be in here

var (
	ErrMessageNotFound = errors.New("Message Not Found")
	ErrMessageInvalid  = errors.New("Message Invalid")
)

type messageService struct {
	messageRepo MessageRepository
}

func NewMessageService(messageRepo MessageRepository) MessageService {
	return &messageService{
		messageRepo,
	}
}

func (m *messageService) DirectMessagesList(userID, receiverID, channelID int64) ([]*Message, error) {
	return m.messageRepo.DirectMessagesList(userID, receiverID, channelID)
}
