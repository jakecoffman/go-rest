package db

import (
	"context"
	"embed"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"strings"
)

var DB *sqlx.DB

func Connect() {
	var err error
	DB, err = sqlx.Connect("postgres", "postgres://localhost:5432/pgx_ex?sslmode=disable")
	if err != nil {
		log.Fatalln("Unable to connect to database:", err)
	}
}

func Migrate() {
	dir, err := migrations.ReadDir("migrations")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range dir {
		if file.IsDir() {
			continue
		}
		data, err := migrations.ReadFile("migrations/" + file.Name())
		if err != nil {
			log.Panicf("Failed reading file %v: %v", file.Name(), err.Error())
		}
		for _, line := range strings.Split(string(data), ";") {
			_, err = DB.ExecContext(context.Background(), line)
			if err != nil {
				log.Panicf("Failed executing sql in file %v: %v", file.Name(), err.Error())
			}
		}
	}
}

//go:embed migrations/*
var migrations embed.FS
