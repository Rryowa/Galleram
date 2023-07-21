package models

import (
	"database/sql"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

type User struct {
	ID           uint
	Email        string
	PasswordHash string
}

type UserService struct {
	DB *sql.DB
}

type NewUser struct {
	Email    string
	Password string
}

func (us *UserService) Create(nu *NewUser) (*User, error) {
	email := strings.ToLower(nu.Email)
	hashBytes, err := bcrypt.GenerateFromPassword([]byte(nu.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("create user: %w", err)
	}
	passwordHash := string(hashBytes)

	//user to return
	user := User{
		Email:        email,
		PasswordHash: passwordHash,
	}
	row := us.DB.QueryRow(`
		INSERT INTO users (email, passwordHash)
		VALUES ($1, $2) RETURNING id`, email, passwordHash)
	err = row.Scan(&user.ID)

	return &user, nil
}
