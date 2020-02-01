//Package arango ...
package arango

import (
	"github.com/amiraliio/tgbp-api/config"
	"github.com/amiraliio/tgbp-api/domain/message"
)

type messageRepo struct{
	appConfig config.AppConfig
}


func NewArangoMessageRepository(appConfig config.AppConfig) message.MessageRepository{
	return &messageRepo{
		appConfig
	}
}


func (a *messageRepo) DirectMessagesList(userID, receiverID, channelID int64) ([]*message.Message, error) {
	panic("implement me")
}