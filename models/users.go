package models

import (
	"context"
	"ewallet_be/utils"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type User struct {
	ID             int    `json:"idUser" db:"id_user"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	PIN            string `json:"pin"`
	Username       string `json:"username"`
	Phone          string `json:"phone"`
	ProfilePicture string `json:"profilePicture" db:"profile_picture"`
}

type UserListItem struct {
	IDUser         int     `json:"idUser" db:"id_user"`
	Username       string  `json:"username"`
	Phone          *string `json:"phone"`
	ProfilePicture *string `json:"profilePicture" db:"profile_picture"`
}

func Register(user User) (int, error) {
	conn, err := utils.ConnectDB()
	if err != nil {
		return 0, err
	}
	defer conn.Release()

	var userID int
	err = conn.QueryRow(
		context.Background(),
		`
    INSERT INTO users (email, password, pin, username, phone, profile_picture)
    VALUES ($1, $2, $3, $4, $5, $6)
    RETURNING id_user
    `,
		user.Email, user.Password, user.PIN, user.Username, user.Phone, user.ProfilePicture,
	).Scan(&userID)

	if err != nil {
		return 0, err
	}
	return userID, nil
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

func FindOneUserByID(id int) (User, error) {
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
    WHERE id_user = $1
    `, id,
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


func EditUser(id int, user User) error {
	conn, err := utils.ConnectDB()
	if err != nil {
		return err
	}
	defer conn.Release()

	result, err := conn.Exec(
		context.Background(),
		`UPDATE users SET email = $1, password = $2, username = $3, phone = $4, profile_picture = $5 WHERE id_user = $6`,
		user.Email, user.Password, user.Username, user.Phone, user.ProfilePicture, id,
	)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return fmt.Errorf("user not found")
	}
	return nil
}

func GetOtherUsers(idUser int, keyword string) ([]UserListItem, error) {
	conn, err := utils.ConnectDB()
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	query := `
		SELECT id_user, username, phone, profile_picture
		FROM users
		WHERE id_user != $1
	`
	args := []any{idUser}

	if keyword != "" {
		query += " AND (username ILIKE $2 OR phone ILIKE $2)"
		args = append(args, "%"+keyword+"%")
	}

	query += " ORDER BY username ASC"

	rows, err := conn.Query(context.Background(), query, args...)
	if err != nil {
		return nil, err
	}

	users, err := pgx.CollectRows(rows, pgx.RowToStructByName[UserListItem])
	if err != nil {
		return nil, err
	}

	return users, nil
}
