package users

import "github.com/jakecoffman/crud"

const root = "/users"

var tags = []string{"Users"}

var Routes = []crud.Spec{{
	Method:  "GET",
	Path:    root,
	Handler: List,
	Tags:    tags,
	Summary: "Lists users",
}, {
	Method:  "POST",
	Path:    root,
	Handler: Add,
	Tags:    tags,
	Summary: "Add user",
	Validate: crud.Validate{
		Body: crud.Object(map[string]crud.Field{
			"name": crud.String().Required().Example("Bob"),
		}),
	},
}, {
	Method:  "GET",
	Path:    root + "/{id}",
	Handler: Get,
	Tags:    tags,
	Summary: "Get user by ID",
	Validate: crud.Validate{
		Path: crud.Object(map[string]crud.Field{
			"id": crud.Integer().Required(),
		}),
	},
}}
