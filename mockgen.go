package main

//go:generate mockgen -package=mock -self_package=freggy.dev -destination=mock/service.go freggy.dev/stats/rpc/go/service StatsService
//go:generate mockgen -package=mock -self_package=freggy.dev -destination=mock/flash.go freggy.dev/stats/pkg/flash DataAccess
