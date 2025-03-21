package main

import "flag"

var GlobalAppFlag AppFlags

type AppFlags struct {
	// Server
	Address string

	// Database
	DB DBFlags
}

type DBFlags struct {
	User     string
	Password string
	Database string
}

func LoadAppFlags() error {
	flag.StringVar(&GlobalAppFlag.Address, "address", ":8080", "Server address")
	flag.StringVar(&GlobalAppFlag.DB.User, "db-user", "", "Database user")
	flag.StringVar(&GlobalAppFlag.DB.Password, "db-password", "", "Database password")
	flag.StringVar(&GlobalAppFlag.DB.Database, "db-database", "", "Database database")
	flag.Parse()
	return nil
}
