package model

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id         int64     `json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Username   string    `json:"username"`
	Password   string    `json:"-"`
	Created_At time.Time `json:"created_at"`
	Updated_At time.Time `json:"-"`
}

func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	u.Password = string(hashedPassword)

	return nil
}
