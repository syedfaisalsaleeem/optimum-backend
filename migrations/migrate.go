package main

import (
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	user := goDotEnvVariable("DB_USER")
	password := goDotEnvVariable("DB_PASSWORD")
	host := goDotEnvVariable("DB_HOST")
	name := goDotEnvVariable("DB_NAME")
	m, err := migrate.New(
		"file://migration",
		"postgres://"+user+":"+password+"@"+host+":5432/"+name+"?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil {
		log.Fatal(err)
	}
}
