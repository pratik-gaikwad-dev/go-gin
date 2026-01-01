package models

import (
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
