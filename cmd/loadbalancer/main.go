package main

import (
	"flag"

	"github.com/jroderic-07/load-balancer/internal/config"
	"github.com/jroderic-07/load-balancer/internal/backend"
	"github.com/jroderic-07/load-balancer/internal/loadbalancer"
)

func main() {
	var port = flag.String("port", "", "port to run reverse proxy on")
	var configFilePath = flag.String("config-file", "", "path to configuration file")
	var loadBalancingType = flag.String("lb-type", "", "type of load balancing")

	flag.Parse()

	configuration := config.New(*configFilePath)
	backends := backend.NewBackendCollection(configuration)

	if *loadBalancingType == "roundRobin" {
		lb := loadbalancer.NewRoundRobin(*port, backends)
		lb.LoadBalancer.Serve(lb)
	}
}
