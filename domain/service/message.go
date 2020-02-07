package service

import (
	"errors"

	"github.com/amiraliio/tgbp-api/domain/model"
)

//TODO repo <-via aport<- service -> serializer -> rest
//TODO validate of data should be in here

var (
	ErrMessageNotFound = errors.New("Message Not Found")
	ErrMessageInvalid  = errors.New("Message Invalid")
)

type MessageService interface {
	DirectMessagesList(userID, receiverID, channelID int64) ([]*model.Message, error)
}

type MessageRepository interface {
	DirectMessagesList(userID, receiverID, channelID int64) ([]*model.Message, error)
}

type MessageSerializer interface {
	Encode(input *model.Message) ([]byte, error)
	Decode(input []byte) (*model.Message, error)
}

type messageService struct {
	messageRepo MessageRepository
}

func NewMessageService(messageRepo MessageRepository) MessageService {
	return &messageService{
		messageRepo,
	}
}

func (m *messageService) DirectMessagesList(userID, receiverID, channelID int64) ([]*model.Message, error) {
	return m.messageRepo.DirectMessagesList(userID, receiverID, channelID)
}
