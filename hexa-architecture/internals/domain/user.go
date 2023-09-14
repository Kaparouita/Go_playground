package domain

import "time"

type User struct {
	Id        uint `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

func NewUser(id uint, firstname, lastname string) *User {
	return &(User{
		Id:        id,
		FirstName: firstname,
		LastName:  lastname,
	})
}
