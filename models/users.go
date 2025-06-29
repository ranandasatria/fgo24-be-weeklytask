package models

import (
	"context"
	"ewallet_be/utils"

	"github.com/jackc/pgx/v5"
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
	INSERT INTO users (email, password, pin, username, phone, profile_picture)
	VALUES ($1, $2, $3, $4, $5, $6)
	`,
		user.Email, user.Password, user.PIN, user.Username, user.Phone, user.ProfilePicture,
	)
	return err
}

func FindOneUserByEmail(email string) (User, error) {
  conn, err := utils.ConnectDB()
  if err != nil {
    return User{}, err
  }
  defer conn.Release()

  rows, err := conn.Query(
    context.Background(),
    `
    SELECT id_user, email, password, pin, username, phone, profile_picture
    FROM users
    WHERE email = $1
    `,
    email,
  )
  if err != nil {
    return User{}, err
  }

  user, err := pgx.CollectOneRow[User](rows, pgx.RowToStructByName)
  if err != nil {
    return User{}, err
  }

  return user, nil
}