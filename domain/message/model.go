package message

import (
	"github.com/amiraliio/tgbp-api/domain/channel"
	"github.com/amiraliio/tgbp-api/domain/user"
)

//TODO validation and msgpack support

type Message struct {
	ID               int64            `json:"id"`
	CreatedAt        string           `json:"createdAt"`
	UpdatedAt        string           `json:"updatedAt"`
	UserID           int64            `json:"userID"`
	ChannelID        int64            `json:"channelID"`
	ParentID         int64            `json:"parentID"`
	ChannelMessageID string           `json:"channelMessageID"`
	BotMessageID     string           `json:"botMessageID"`
	Message          string           `json:"message"`
	Receiver         int64            `json:"receiver"`
	Type             string           `json:"type"`
	User             *user.User       `json:"user"`
	Channel          *channel.Channel `json:"channel"`
}
