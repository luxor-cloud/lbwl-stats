package main

import (
	"flag"
)

func main() {
	flag.String("config", "/etc/lbwl-stats/config.json", "Config")
	flag.Parse()

}
