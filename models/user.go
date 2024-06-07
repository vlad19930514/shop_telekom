package models

import (
	"context"
	"errors"

	"github.com/vlad19930514/shop_telekom/db"
	"github.com/vlad19930514/shop_telekom/utils"
)

type User struct {
	ID                int64
	Email             string
	Password          string
	Username          string
	Is_email_verified bool
}

func (u User) Save(ctx context.Context) error {

	const createUser = `INSERT INTO users
		(email, password)
		VALUES($1, $2);`

	hashedPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}
	_, err = db.ConnPool.Exec(ctx, createUser, u.Email, hashedPassword)
	if err != nil {
		return err
	}
	return nil
}
func (u *User) ValidateCredentials(ctx context.Context) error {
	query := "SELECT email, password, id, username FROM users WHERE email=$1"

	row := db.ConnPool.QueryRow(ctx, query, u.Email)

	var retrievedPassword string
	err := row.Scan(&u.Email, &retrievedPassword, &u.ID, &u.Username)

	if err != nil {
		return errors.New("credentials invalid")
	}

	passwordIsValid := utils.CheckPasswordHash(u.Password, retrievedPassword)

	if !passwordIsValid {
		return errors.New("credentials invalid")
	}

	return nil
}
func (u User) ChangeUsername(ctx context.Context) error {

	query := "update users set username=$1 where id=$2"

	_, err := db.ConnPool.Exec(ctx, query, u.Username, u.ID)

	if err != nil {
		return errors.New("can't update username ")
	}

	return nil
}
