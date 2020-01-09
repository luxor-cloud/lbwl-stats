package main

import (
	"flag"
	"freggy.dev/stats/internal/statsserver"
	"freggy.dev/stats/pkg/flash"
	"freggy.dev/stats/rpc/go/service"
	"log"
	"net/http"
)

func main() {
	path := flag.String("config", "/etc/lbwl-stats/config.json", "Config")
	flag.Parse()

	config, err := statsserver.ReadConfig(*path)
	if err != nil {
		log.Fatalf("%v", err)
	}

	var flashDAO flash.DataAccess

	if config.FlashDB["type"] == "sql" {
		conn, err := flash.ConnectSQL(config.FlashDB["Address"])
		if err != nil {
			log.Fatalf("%v", err)
		}

		defer conn.Close()
		flashDAO = &flash.SQLDataAccess{Conn: conn}
	}

	handler := &statsserver.Server{FlashDAO: flashDAO}
	server := service.NewStatsServiceServer(handler, nil)

	if config.UseTLS {
		log.Fatal(http.ListenAndServeTLS(config.Address, config.TLSCert, config.TLSKey, server))
	} else {
		log.Fatal(http.ListenAndServe(config.Address, server))
	}
}
