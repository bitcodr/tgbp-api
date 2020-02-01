package user

import (
	"github.com/amiraliio/tgbp-api/domain/channel"
)

type User struct {
	ID        int64         `json:"id"`
	Status    string        `json:"status"`
	UserID    string        `json:"userId"`
	Username  string        `json:"username"`
	FirstName string        `json:"firstName"`
	LastName  string        `json:"lastName"`
	Lang      string        `json:"lang"`
	Email     string        `json:"email"`
	IsBot     string        `json:"isBot"`
	CustomID  string        `json:"customID"`
	CreatedAt string        `json:"createdAt"`
	UpdatedAt string        `json:"updatedAt"`
	UserSign  *UserUserName `json:"userSign"`
}

type UserUserName struct {
	ID        int64            `json:"id"`
	UserID    int64            `json:"userID"`
	ChannelID int64            `json:"channelID"`
	Username  string           `json:"username"`
	CreatedAt string           `json:"createdAt"`
	UpdatedAt string           `json:"updatedAt"`
	User      *User            `json:"user"`
	Channel   *channel.Channel `json:"channel"`
}
