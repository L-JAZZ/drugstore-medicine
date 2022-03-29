package database

import (
	"log"
	"medicine/utils/env"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var db *sqlx.DB

// DB - Singleton Database connection
func DB() *sqlx.DB {

	if db == nil {

		newDb, err := sqlx.Connect(env.EnvDBDriver(), env.EnvServerConnString())

		if err != nil {
			log.Fatal(err)
		}

		db = newDb
	}
	return db
}
