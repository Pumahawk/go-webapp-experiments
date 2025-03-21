package main

import "flag"

var GlobalAppFlag AppFlags

type AppFlags struct {
	address string
}

func LoadAppFlags() error {
	flag.StringVar(&GlobalAppFlag.address, "address", ":8080", "Server address")
	flag.Parse()
	return nil
}
