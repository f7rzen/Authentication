package main

import (
	"database/sql"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func CreateUser(db *sql.DB, user User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("unable to hash password: %v", err)
	}

	query := `
    INSERT INTO users (username, email, password)
    VALUES ($1, $2, $3)`

	_, err = db.Exec(query, user.Name, user.Email, string(hashedPassword))
	if err != nil {
		return fmt.Errorf("unable to create user: %v", err)
	}
	return nil
}

func GetUser(db *sql.DB, email string) (*User, error) {
	query := `
    SELECT username, password
    FROM users	
    WHERE email = $1`

	var user User
	row := db.QueryRow(query, email)
	err := row.Scan(&user.Name, &user.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("unable to get user: %v", err)
	}
	return &user, nil
}
