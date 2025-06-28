package models

import (
	"context"
	"ewallet_be/utils"
)

type User struct {
	ID             int    `json:"iduser" db:"id_user"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	PIN            string `json:"pin"`
	Username       string `json:"username"`
	Phone          string `json:"phone"`
	ProfilePicture string `json:"profilepicture" db:"profile_picture"`
}

func Register(user User) error {
	conn, err := utils.ConnectDB()
	if err != nil {
		return err
	}

	defer conn.Release()

	_, err = conn.Exec(
		context.Background(),
		`
		INSERT INTO users (email, password, pin, username)
		VALUES
		($1, $2, $3, $4)
		`,
		user.Email, user.Password, user.PIN, user.Username,
	)
	return err
}
