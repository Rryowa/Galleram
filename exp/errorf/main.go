package main

import (
	"errors"
	"fmt"
)

func Connect() error {
	return errors.New("conn failed")
}

func CreateUser() error {
	err := Connect()
	if err != nil {
		return fmt.Errorf("create user -> %w", err)
	}
	return nil
}

func CreateOrg() error {
	err := CreateUser()
	if err != nil {
		return fmt.Errorf("create org -> %w", err)
	}
	return nil
}

func main() {
	err := CreateOrg()
	fmt.Println(err)
}
