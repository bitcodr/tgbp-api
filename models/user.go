//Package models ...
package models

type User struct {
	ID        int64      `json:"id"`
	Status    string     `json:"status"`
	UserID    string     `json:"userId"`
	Username  string     `json:"username"`
	FirstName string     `json:"firstName"`
	LastName  string     `json:"lastName"`
	Lang      string     `json:"lang"`
	Email     string     `json:"email"`
	IsBot     string     `json:"isBot"`
	CustomID  string     `json:"customID"`
	CreatedAt string     `json:"createdAt"`
	UpdatedAt string     `json:"updatedAt"`
}