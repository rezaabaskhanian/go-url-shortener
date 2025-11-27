package acesscontrol

import (
	"game_app/internal/repository/mysql"
)

type DB struct {
	conn *mysql.MySqlDB
}

func New(conn *mysql.MySqlDB) *DB {
	return &DB{conn: conn}
}
