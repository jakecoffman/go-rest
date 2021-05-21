package test

import (
	"fmt"
	"github.com/jakecoffman/rest/lib/db"
	"github.com/jmoiron/sqlx"
	"github.com/ory/dockertest/v3"
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	log.SetFlags(log.Lshortfile)

	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	// pulls an image, creates a container based on it and runs it
	log.Println("pulling docker container")
	resource, err := pool.Run("postgres", "13", []string{
		"POSTGRES_USER=test",
		"POSTGRES_PASSWORD=test",
		"listen_addresses = '*'",
	})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	log.Println("connecting to DB")
	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	if err = pool.Retry(func() error {
		var err error
		db.DB, err = sqlx.Open("postgres", fmt.Sprintf("postgres://test:test@localhost:%s/postgres?sslmode=disable", resource.GetPort("5432/tcp")))
		if err != nil {
			return err
		}
		// connection errors? uncomment this line to reveal them
		//log.Println(db.DB.Ping())
		return db.DB.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}
	db.Migrate()

	code := m.Run()

	// You can't defer this because os.Exit doesn't care for defer
	if err = pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	os.Exit(code)
}
