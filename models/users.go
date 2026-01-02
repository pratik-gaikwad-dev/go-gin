package models

import (
	"errors"

	"github.com/learn-gin/db"
	"github.com/learn-gin/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (user User) Save() error {
	query := `
	INSERT INTO users (email, password)
	VALUES (?, ?)
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	hashPass, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(user.Email, hashPass)
	if err != nil {
		return err
	}
	return nil
}

func (user *User) ValidateCredentials() error {
	query := `SELECT id, password FROM users WHERE email = ?`
	row := db.DB.QueryRow(query, user.Email)

	var retriveString string
	err := row.Scan(&user.ID, &retriveString)
	if err != nil {
		return errors.New("Invalid Credentials!")
	}

	passwordIsValid := utils.CheckPassword(user.Password, retriveString)

	if !passwordIsValid {
		return errors.New("Invalid Credentials!")
	}
	return nil
}
