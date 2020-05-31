package main

//go:generate mockgen -package=mock -self_package=freggy.dev -destination=mock/service.go freggy.dev/stats/service/v1 StatsServiceServer
//go:generate mockgen -package=mock -self_package=freggy.dev -destination=mock/flash.go freggy.dev/stats/pkg/flash DataAccess
