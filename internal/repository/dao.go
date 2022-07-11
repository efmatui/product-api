package repository

import (
	"github.com/go-pg/pg"
)

func NewDB() *pg.DB {
	db := pg.Connect(&pg.Options{
		User:     "sellers",
		Password: "",
		Database: "postgres",
		Addr:     "127.0.0.1:5432",
	})
	return db
}
