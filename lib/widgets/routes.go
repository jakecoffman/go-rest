package widgets

import "github.com/jakecoffman/crud"

const root = "/widgets"

var tags = []string{"Widgets"}

var Routes = []crud.Spec{{
	Method:  "GET",
	Path:    root,
	Handler: List,
	Tags:    tags,
	Summary: "Lists widgets",
}, {
	Method:  "POST",
	Path:    root,
	Handler: Add,
	Tags:    tags,
	Summary: "Add widget",
	Validate: crud.Validate{
		Body: crud.Object(map[string]crud.Field{
			"type":  crud.String().Required().Example("curly"),
			"owner": crud.Integer().Example(1).Description("ID of a user"),
		}),
	},
}}
