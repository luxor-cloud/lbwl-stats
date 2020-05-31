package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"freggy.dev/stats/pkg/flash"
	servicev1 "freggy.dev/stats/service/v1"
	"freggy.dev/stats/server"
	serverv1 "freggy.dev/stats/server/v1"
)

func main() {
	/*
	path := flag.String("config", "/etc/lbwl-stats/config.json", "Config")
	flag.Parse()

	config, err := statsserver.ReadConfig(*path)
	if err != nil {
		log.Fatalf("%v", err)
	}*/


	config := server.DefaultConfig

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

	listener, err := net.Listen("tcp", ":1337")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	servicev1.RegisterStatsServiceServer(grpcServer, serverv1.NewServer(conn))
	log.Fatal(grpcServer.Serve(listener))
}
