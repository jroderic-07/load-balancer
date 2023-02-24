package main

import (
	"flag"

	"github.com/jroderic-07/load-balancer/internal/config"
	"github.com/jroderic-07/load-balancer/internal/backend"
	"github.com/jroderic-07/load-balancer/internal/loadbalancer"
)

func main() {
	var port = flag.Int("port", 0, "port to run reverse proxy on")
	var configFilePath = flag.String("config-file", "", "path to configuration file")
	var loadBalancingType = flag.String("lb-type", "", "type of load balancing")

	flag.Parse()

	test := config.New(*configFilePath)
	test1 := backend.NewBackendCollection(test)

	if *loadBalancingType == "roundRobin" {
		test2 := loadbalancer.NewRoundRobin(*port, test1)
		test2.LoadBalancer.Serve(test2)
	}
}
