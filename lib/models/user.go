package models

import (
	"context"
	"github.com/jakecoffman/rest/lib/db"
	"log"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func FindUsers(c context.Context) ([]User, error) {
	results := []User{}
	err := db.DB.SelectContext(c, &results, `select * from "USER"`)
	return results, err
}

func GetUserById(c context.Context, id int) (*User, error) {
	var u User
	err := db.DB.GetContext(c, &u, `select * from "USER" where id=$1`, id)
	return &u, err
}

func InsertUser(c context.Context, payload *User) error {
	stmt, err := db.DB.PrepareNamedContext(c, `insert into "USER" (name) values(:name) returning *`)
	if err != nil {
		log.Println(err)
		return err
	}

	err = stmt.Get(payload, payload)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
