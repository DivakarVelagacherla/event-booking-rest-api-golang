package models

import (
	"event-booking-rest-api-golang/database"
	"event-booking-rest-api-golang/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error {
	query := `
		INSERT INTO users (email, password)
		VALUES (?,?)
	`
	stmt, err := database.DB.Prepare(query)
	if err != nil {
		return err
	}

	hashedPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}

	_, err = stmt.Exec(u.Email, hashedPassword)
	if err != nil {
		return err
	}

	return err
}
