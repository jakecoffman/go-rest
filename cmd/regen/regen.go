package main

import (
	"github.com/jakecoffman/rest/lib/db"
	"log"
	"os"
	"os/exec"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	db.Connect("postgres://postgres:@127.0.0.1:5432?sslmode=disable")
	if _, err := db.DB.Exec("drop database if exists go_rest_test"); err != nil {
		log.Fatal(err)
	}
	if _, err := db.DB.Exec("create database go_rest_test"); err != nil {
		log.Fatal(err)
	}
	if err := db.DB.Close(); err != nil {
		log.Fatal(err)
	}
	_ = db.DB.Close()
	db.Connect("postgres://postgres:@127.0.0.1:5432/go_rest_test?sslmode=disable")

	db.Migrate()
	_ = db.DB.Close()

	cmd := exec.Command("sqlboiler", "--struct-tag-casing=camel", "--no-tests", "--wipe", "psql")
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	if cmd.Run() != nil {
		log.Fatal("sqlboiler failed, see output above")
	}
}
