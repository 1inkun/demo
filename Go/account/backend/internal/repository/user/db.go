package user

import "database/sql"

type Queries struct {
	db *sql.DB
}

func NewQueries (db *sql.DB) *Queries{
	return &Queries{db : db}
}
