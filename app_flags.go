package main

import "flag"

var GlobalAppFlag AppFlags

type AppFlags struct {
	// Server
	Address string
}

func LoadAppFlags() error {
	flag.StringVar(&GlobalAppFlag.Address, "address", ":8080", "Server address")
	flag.Parse()
	return nil
}
