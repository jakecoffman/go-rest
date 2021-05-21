package models

import (
	"context"
	"github.com/jakecoffman/rest/lib/db"
	"log"
)

type Widget struct {
	ID      int    `json:"id"`
	Type    string `json:"type"`
	OwnerID int    `json:"ownerID" db:"owner_id"`

	Owner *User `json:"owner,omitempty" db:"-"`
}

func FindWidgets(c context.Context) ([]Widget, error) {
	results := []Widget{}
	err := db.DB.SelectContext(c, &results, `select * from widget`)
	return results, err
}

func GetWidgetById(c context.Context, id int) (*Widget, error) {
	var instance Widget
	err := db.DB.GetContext(c, &instance, `select * from widget where id=$1`, id)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &instance, err
}

func InsertWidget(c context.Context, payload *Widget) error {
	stmt, err := db.DB.PrepareNamedContext(c, `insert into widget (type, owner_id) values(:type, :owner_id) returning *`)
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
