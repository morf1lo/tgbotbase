package postgres

import "github.com/jackc/pgx/v5"

type Postgres struct {
}

func New(db *pgx.Conn) *Postgres {
	return &Postgres{}
}
