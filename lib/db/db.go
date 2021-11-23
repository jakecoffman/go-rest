package db

import (
	"context"
	"database/sql"
	"embed"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"log"
	"strings"
	"time"
)

var DB *sql.DB

const (
	defaultMaxConnLifetime = time.Hour
	defaultMaxConnIdleTime = time.Minute * 30
)

func Connect(db string) {
	var err error
	DB, err = sql.Open("pgx", db)
	if err != nil {
		log.Fatal(err)
	}
	DB.SetConnMaxIdleTime(defaultMaxConnIdleTime)
	DB.SetConnMaxLifetime(defaultMaxConnLifetime)
	DB.SetMaxOpenConns(50)
	boil.SetDB(DB)
}

func Migrate() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

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
			log.Fatalf("Failed reading file %v: %v", file.Name(), err.Error())
		}
		for _, line := range strings.Split(string(data), ";") {
			_, err := DB.ExecContext(ctx, line)
			if err != nil {
				log.Fatalf("Failed executing sql in file %v: %v", file.Name(), err.Error())
			}
		}
	}
}

//go:embed migrations/*
var migrations embed.FS
