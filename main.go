package main

import (
	"log"

	"freggy.dev/stats/internal/statsserver"
	"freggy.dev/stats/pkg/flash"
)

func main() {
	/*
	path := flag.String("config", "/etc/lbwl-stats/config.json", "Config")
	flag.Parse()

	config, err := statsserver.ReadConfig(*path)
	if err != nil {
		log.Fatalf("%v", err)
	}*/


	config := statsserver.DefaultConfig

	conn, err := flash.ConnectSQL(
		config.FlashDBConnection.Username,
		config.FlashDBConnection.Password,
		config.FlashDBConnection.Host,
		config.FlashDBConnection.Database,
	)

	if err != nil {
		log.Fatalf("%v", err)
	}

	defer conn.Close(nil)
}
