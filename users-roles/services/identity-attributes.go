package services

import (
	"context"
	"fmt"
	"simpl-go/users-roles/server"
	"strings"
)

func IdentityAttributeSearch(ctx context.Context) ([]IdentityAttributeInfo, error) {
	conn, err := server.DBConn(ctx)
	if err != nil {
		return nil, fmt.Errorf("Error retrieve connection for role identity attribute search. %w", err)
	}

	sb := strings.Builder{}
	sb.WriteString(`
		select 
			id
		from
			identity_attribute
		where
			1=1
	`)
	query := sb.String()

	rows, err := conn.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("Unable to retrieve role from database. %w", err)
	}

	var idas []IdentityAttributeInfo
	for ;rows.Next(); {
		var ida IdentityAttributeInfo
		if err = rows.Scan(&ida.Id); err != nil{
			return nil, fmt.Errorf("Unable to scan identity attribute: %w", err)
		}
		idas = append(idas, ida)

	}

	return idas, nil
}
