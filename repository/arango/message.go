//Package arango ...
package arango

import (
	"github.com/amiraliio/tgbp-api/domain/model"
	"github.com/amiraliio/tgbp-api/domain/service"
	"github.com/amiraliio/tgbp-api/config"
)

type messageRepo struct {
	appConfig *config.App
}

func NewArangoMessageRepository(appConfig *config.App) service.MessageRepository {
	return &messageRepo{
		appConfig,
	}
}

func (a *messageRepo) DirectMessagesList(userID, receiverID, channelID int64) ([]*model.Message, error) {
	panic("implement me")
}
