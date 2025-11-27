package acesscontrol

import (
	"game_app/internal/entity"
	"game_app/internal/pkg/slice"
	"game_app/internal/repository/mysql"
	"strings"
)

func (d *DB) GetUserPermissionTitle(userID uint, role entity.Role) ([]entity.PermissionTitle, error) {

	rolAcl := make([]entity.AccessControl, 0)

	rows, err := d.conn.Conn().Query(`select * from access_controls where actor_type=? and actor_id = ?`,
		entity.RoleActorType, role)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	//var acl entity.AccessControl
	for rows.Next() {
		acl, err := scanAccessControl(rows)

		if err != nil {
			return nil, err
		}

		rolAcl = append(rolAcl, acl)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	userAcl := make([]entity.AccessControl, 0)
	userRows, err := d.conn.Conn().Query(`select * from access_controls where actor_type=? and actor_id = ?`,
		entity.UserActorType, userID)

	if err != nil {
		return nil, err
	}

	defer userRows.Close()

	//var acl entity.AccessControl
	for userRows.Next() {
		user, err := scanAccessControl(userRows)

		if err != nil {
			return nil, err
		}

		userAcl = append(userAcl, user)
	}

	if err := userRows.Err(); err != nil {
		return nil, err
	}

	permissionIDs := make([]uint, 0)

	for _, r := range rolAcl {
		if !slice.DoesExist(permissionIDs, r.PermissionID) {
			permissionIDs = append(permissionIDs, r.PermissionID)
		}

	}

	if len(permissionIDs) == 0 {
		return nil, nil
	}

	args := make([]any, len(permissionIDs))

	for i, id := range permissionIDs {
		args[i] = id
	}

	query := "select * from permissions where where id in (?" + strings.Repeat(",?", len(permissionIDs)-1) + ")"

	pRows, err := d.conn.Conn().Query(query, args...)

	if err != nil {
		return nil, err
	}

	defer pRows.Close()

	permissionTitles := make([]entity.PermissionTitle, 0)

	for pRows.Next() {
		permission, err := scanPermission(pRows)
		if err != nil {
			return nil, err
		}

		permissionTitles = append(permissionTitles, permission.Title)
	}

	if err := rows.Err(); err != nil {

		return nil, err

	}
	return permissionTitles, nil
}

func scanAccessControl(scanner mysql.Scanner) (entity.AccessControl, error) {
	acl := entity.AccessControl{}

	var createdUser []uint8

	err := scanner.Scan(&acl.ID, &acl.ActorID, &acl.ActorType, &acl.PermissionID, &createdUser)

	return acl, err

}
