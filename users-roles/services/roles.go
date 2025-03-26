package services

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"simpl-go/users-roles/server"
)

func RoleFindById(ctx context.Context, id string) (*RoleInfo, error) {
	conn, err := server.DBConn(ctx)
	if err != nil {
		return nil, fmt.Errorf("Error retrieve connection for role use. %w", err)
	}
	var roleInfo RoleInfo
	query :=  "select id from roles where id = $1"
	row := conn.QueryRowContext(ctx, query, id)
	err = row.Scan(&roleInfo.Id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		} else {
			log.Printf("Unable to retrieve role from database. %s. Query: %s", query, err)
			return nil, fmt.Errorf("Unable to retrieve role from database. %w", err)
		}
	} else {
		return &roleInfo, nil
	}
}
