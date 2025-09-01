package models

import (
	"event-booking-rest-api-golang/database"
)

type Registration struct {
	EventId int64
	UserId  int64
}

func (r *Registration) Save() error {
	query := `
		INSERT INTO registrations (event_id, user_id)
		VALUES (?,?)
	`

	stmt, err := database.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(r.EventId, r.UserId)
	if err != nil {
		return err
	}

	return nil
}

func (r *Registration) Cancel() error {
	query := `
		DELETE FROM registrations
		WHERE event_id = ? AND user_id = ?
	`

	stmt, err := database.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(r.EventId, r.UserId)
	if err != nil {
		return err
	}

	return nil

}
