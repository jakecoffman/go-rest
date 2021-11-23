package authors

import "github.com/jakecoffman/crud"

const root = "/authors"

var tags = []string{"Authors"}

var Routes = []crud.Spec{
	{
		Method:  "GET",
		Path:    root,
		Handler: list,
		Validate: crud.Validate{
			Query: crud.Object(map[string]crud.Field{
				"limit":  crud.Integer().Min(0).Default(100),
				"offset": crud.Integer().Min(0).Default(0),
				"sort":   crud.String().Default("created"),
				"order":  crud.String().Enum("asc", "desc").Default("desc"),
			}),
		},
		Tags:    tags,
		Summary: "Lists authors",
	}, {
		Method:  "POST",
		Path:    root,
		Handler: add,
		Tags:    tags,
		Summary: "Add author",
		Validate: crud.Validate{
			Body: crud.Object(map[string]crud.Field{
				"name": crud.String().Required().Example("Bob"),
				"born": crud.DateTime(),
			}),
		},
	}, {
		Method:  "GET",
		Path:    root + "/{id}",
		Handler: get,
		Tags:    tags,
		Summary: "Get user by ID",
		Validate: crud.Validate{
			Path: crud.Object(map[string]crud.Field{
				"id": crud.Integer().Required(),
			}),
		},
	}, {
		Method:  "PATCH",
		Path:    root + "/{id}",
		Handler: patch,
		Tags:    tags,
		Summary: "Update author by ID",
		Validate: crud.Validate{
			Path: crud.Object(map[string]crud.Field{
				"id": crud.Integer().Required(),
			}),
			Body: crud.Object(map[string]crud.Field{
				"name": crud.String(),
				"born": crud.DateTime(),
			}),
		},
	}, {
		Method:  "DELETE",
		Path:    root + "/{id}",
		Handler: remove,
		Tags:    tags,
		Summary: "Delete user by ID",
		Validate: crud.Validate{
			Path: crud.Object(map[string]crud.Field{
				"id": crud.Integer().Required(),
			}),
		},
	},
}
