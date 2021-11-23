package main

import (
	"github.com/jakecoffman/crud"
	adapter "github.com/jakecoffman/crud/adapters/gin-adapter"
	"github.com/jakecoffman/rest/lib/db"
	"github.com/jakecoffman/rest/lib/endpoints/authors"
	"log"
)

func main() {
	log.SetFlags(log.Lshortfile)

	db.Connect("postgres://localhost:5432/go_rest?sslmode=disable")
	db.Migrate()

	engine := adapter.New()
	router := crud.NewRouter("My API", "1.0.0", engine)

	check(router.Add(authors.Routes...))
	//check(router.Add(books.Routes...))

	log.Println("Serving on http://127.0.0.1:8999")
	if err := router.Serve("127.0.0.1:8999"); err != nil {
		log.Panicln(err)
	}
}

func check(err error) {
	if err != nil {
		log.Panicln(err)
	}
}
