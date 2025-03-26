package services

import (
	"context"
	"fmt"
	"log"
	"simpl-go/users-roles/db"
	"simpl-go/users-roles/server"
)

func IdentityAttributeSearch(ctx context.Context, params IdentityAttributeSearchParams) ([]IdentityAttributeInfo, error) {
	conn, err := server.DBConn(ctx)
	if err != nil {
		return nil, fmt.Errorf("Error retrieve connection for role identity attribute search. %w", err)
	}

	query, err := db.Query("ida-search.sql", params)
	if err != nil {
		return nil, fmt.Errorf("Unable to retrieve query from template: %w", err)
	}

	log.Printf("Result params: %v", query.Params)
	rows, err := conn.QueryContext(ctx, query.Sql, query.Params...)
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
