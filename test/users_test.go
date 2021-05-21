package test

import (
	"context"
	"github.com/jakecoffman/rest/lib/models"
	"testing"
)

func TestUsers(t *testing.T) {
	c := context.Background()
	user := &models.User{
		Name: "Alice",
	}
	if err := models.InsertUser(c, user); err != nil {
		t.Fatal(err)
	}
	if user.ID == 0 {
		t.Error(user.ID)
	}
	user, err := models.GetUserById(c, user.ID)
	if err != nil {
		t.Fatal(err)
	}
	if user.Name != "Alice" {
		t.Error(user)
	}
}
