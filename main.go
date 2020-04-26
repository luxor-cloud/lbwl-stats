package main

import (
	"log"
	"net/http"

	"freggy.dev/stats/internal/statsserver"
	"freggy.dev/stats/pkg/flash"
	"freggy.dev/stats/rpc/go/service"
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

	defer conn.Close()

	handler := &statsserver.Server{FlashDAO: conn}
	server := service.NewStatsServiceServer(handler, nil)

	if config.UseTLS {
		log.Fatal(http.ListenAndServeTLS(config.Address, config.TLSCert, config.TLSKey, server))
	} else {
		log.Fatal(http.ListenAndServe(config.Address, server))
	}
}
