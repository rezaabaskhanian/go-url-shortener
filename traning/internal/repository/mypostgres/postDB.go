package repository

import (
	"database/sql"
	_ "database/sql"

	_ "github.com/jackc/pgx/v5"
	_ "github.com/lib/pq"
)

type Config struct {
	User    string
	DBname  string
	Sslmode string
}

type PostDB struct {
	db *sql.DB
}

func (p *PostDB) Conn() *sql.DB {
	return p.db
}

func New(config Config) {
	
}
