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
	hashBytes, err := bcrypt.GenerateFromPassword([]byte(nu.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("create user: %w", err)
	}
	passwordHash := string(hashBytes)

	//user to return
	user := User{
		Email:        strings.ToLower(nu.Email),
		PasswordHash: passwordHash,
	}
	row := us.DB.QueryRow(`
		INSERT INTO users (email, password_hash) 
		VALUES ($1, $2) RETURNING id`, user.Email, user.PasswordHash)
	err = row.Scan(&user.ID)
	if err != nil {
		return nil, fmt.Errorf("insert users: %w", err)
	}
	return &user, nil
}

func (us *UserService) Auth(nu *NewUser) (*User, error) {
	//user to find
	user := User{
		Email: strings.ToLower(nu.Email),
	}

	row := us.DB.QueryRow(`
		SELECT id, password_hash FROM users
		WHERE email=$1;`, user.Email)

	//Extract id and hashed password from db
	err := row.Scan(&user.ID, &user.PasswordHash)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("user not found: %w", err)
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(nu.Password))
	if err != nil {
		return nil, fmt.Errorf("wrong password: %w", err)
	}
	return &user, nil
}
