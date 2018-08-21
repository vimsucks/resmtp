package model

import (
	"github.com/jmoiron/sqlx"
	"github.com/vimsucks/resmtp/config"
	"log"
)

var DB *sqlx.DB

func InitDB(conf config.Config) {
	var err error
	DB, err = sqlx.Connect("mysql", conf.Database.URL)
	if err != nil {
		log.Fatalf("error connecting to database: \"%v\"", err)
	}
}
