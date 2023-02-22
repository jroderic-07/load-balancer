package main

import (
	"flag"

	"github.com/jroderic-07/load-balancer/internal/config"
	"github.com/jroderic-07/load-balancer/internal/backend"
)

func main() {
	var configFilePath = flag.String("config-file", "", "path to configuration file")

	flag.Parse()

	test := config.New(*configFilePath)

	test3 := backend.NewBackendCollection(test)
	print(test3.Backends[1].GetURL())
}
