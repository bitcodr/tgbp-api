//Package repository ...
package repository

import (
	"github.com/amiraliio/tgbp-api/config"
	"github.com/amiraliio/tgbp-api/models"
)

type Repo interface {
	GetAllDM(userID, receiverID int64) ([]*models.Message, error)
}

type RepoService struct{}

//TODO pagination for getAll messsages

func (repo *RepoService) GetAllDM(userID, receiverID int64) ([]*models.Message, error) {
	app := new(config.App)
	app = app.SetAppConfig()
	db := app.DB()
	defer db.Close()
	rows, err := db.Query("select me.message, me.createdAt, us.customID from messages as me inner join users as us on me.userID=us.userID where me.type=? and ((me.userID=? and me.receiver=?) or (me.receiver=? and me.userID=?)) order by me.createdAt asc", "DM", userID, receiverID, userID, receiverID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var messages []*models.Message
	for rows.Next() {
		message := new(models.Message)
		user := new(models.User)
		if err := rows.Scan(&message.Message, &message.CreatedAt, &user.CustomID); err != nil {
			return nil, err
		}
		message.User = user
		messages = append(messages, message)
	}
	return messages, nil
}
