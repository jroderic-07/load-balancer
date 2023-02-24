package main

import (
	"flag"

	"github.com/jroderic-07/load-balancer/internal/config"
	"github.com/jroderic-07/load-balancer/internal/backend"
	"github.com/jroderic-07/load-balancer/internal/loadbalancer"
)

func main() {
	var configFilePath = flag.String("config-file", "", "path to configuration file")
	var loadBalancingType = flag.String("lb-type", "", "type of load balancing")

	flag.Parse()

	test := config.New(*configFilePath)
	test1 := backend.NewBackendCollection(test)
	print(test1.Backends[1].GetURL())

	if *loadBalancingType == "roundRobin" {
		test2 := loadbalancer.NewRoundRobin()
		test2.LoadBalancer.Serve(test2)
	}
}
