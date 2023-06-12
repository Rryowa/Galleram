package main

import (
	"fmt"
	"html/template"
	"os"
)

type User struct {
	Name string
	Age  int
	Meta UserMeta
}

type UserMeta struct {
	Visits int
}

func main() {
	fmt.Println("experimental main")
	t, err := template.ParseFiles("home.gohtml")
	if err != nil {
		panic(err)
	}

	user := User{
		Name: "Jon Doe",
		Age:  111,
		Meta: UserMeta{
			Visits: 4,
		},
	}

	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}
}
