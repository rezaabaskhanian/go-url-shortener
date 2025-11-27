package acesscontrol

import (
	"game_app/internal/entity"
	"game_app/internal/repository/mysql"
	"time"
)

func scanPermission(scanner mysql.Scanner) (entity.Permission, error) {
	var createdAt time.Time
	var p entity.Permission

	err := scanner.Scan(&p.ID, &p.Title, &createdAt)

	return p, err
}
