//Package message ...
package message

import (
	"github.com/pkg/errors"
	"github.com/amiraliio/tgbp-api/config"
	"github.com/amiraliio/tgbp-api/domain/message"
)

// //TODO pagination for getAll messsages


type messageRepo struct{}

func NewMessageRepository() message.MessageRepository {
	return &messageRepo{}
}

func (m *messageRepo) DirectMessagesList(userID, receiverID, channelID int64) (messages []*message.Message, error) {
	app := new(config.App)
	app = app.SetAppConfig()
	db := app.DB()
	defer db.Close()
	rows, err := db.Query("select me.userID, me.message, me.createdAt, uu.username,cha.channelName, cha.channelType from messages as me inner join users as us on me.userID=us.userID inner join users_usernames as uu on uu.userID=us.id and me.channelID=uu.channelID inner join channels as cha on cha.id=me.channelID where me.type=? and me.channelID=? and ((me.userID=? and me.receiver=?) or (me.receiver=? and me.userID=?)) order by me.createdAt asc, me.id asc", "DM", channelID, userID, receiverID, userID, receiverID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		message := new(models.Message)
		user := new(models.User)
		username := new(models.UserUserName)
		channel := new(models.Channel)
		if err := rows.Scan(&message.UserID, &message.Message, &message.CreatedAt, &username.Username, &channel.ChannelName, &channel.ChannelType); err != nil {
			return nil, errors.Wrap(err,"repository.mysql.Message.DirectMessagesList")
		}
		if message.UserID == userID {
			username.Username = "You"
		} else {
			username.Username = "[User " + username.Username + "]"
		}
		user.UserSign = username
		message.User = user
		message.Channel = channel
		messages = append(messages, message)
	}
	return messages, nil
}


