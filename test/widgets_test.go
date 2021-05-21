package test

import (
	"context"
	"github.com/jakecoffman/rest/lib/models"
	"testing"
)

func TestWidgets(t *testing.T) {
	c := context.Background()

	user := &models.User{Name: "Jane"}
	check(models.InsertUser(c, user))

	widget := &models.Widget{
		Type:    "curly",
		OwnerID: user.ID,
	}
	if err := models.InsertWidget(c, widget); err != nil {
		t.Fatal(err)
	}
	if widget.ID == 0 {
		t.Error(widget.ID)
	}

	widget, err := models.GetWidgetById(c, widget.ID)
	if err != nil {
		t.Fatal(err)
	}
	if widget.Type != "curly" {
		t.Error(widget)
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
