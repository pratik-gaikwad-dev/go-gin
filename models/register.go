package models

import (
	"github.com/learn-gin/db"
)

type Register struct {
	ID      int64
	EventID int64
	UserID  int64
}

func (r *Register) Save() error {
	query := `
	INSERT INTO registrations (event_id, user_id)
	VALUES (?, ?)
	`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(r.EventID, r.UserID)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	r.ID = id
	return err
}
