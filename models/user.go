package models

import (
	"errors"
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

func (user *User) ValidateCredentials() error {
	query := "SELECT id, password FROM users WHERE email =?"
	row := database.DB.QueryRow(query, user.Email)

	var retrievedPassword string
	err := row.Scan(&user.ID, &retrievedPassword)
	if err != nil {
		return err
	}

	IsValidCredentials := utils.ComparePasswords(retrievedPassword, user.Password)

	if !IsValidCredentials {
		return errors.New("invalid credentials")
	}

	return nil
}
