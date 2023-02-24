package loadbalancer

import (
	"github.com/jroderic-07/load-balancer/internal/backend"	
)

type loadBalancerInt interface {
	Test()
}

type LoadBalancer struct {
	port int
	backends []*backend.Backend
}

func New(port int, backend []*backend.Backend) *LoadBalancer {
	return &LoadBalancer{port, backend}
}

func (l *LoadBalancer) getPort() int {
	return l.port
}

func (l *LoadBalancer) getBackends() []*backend.Backend {
	return l.backends
}

func (l *LoadBalancer) Serve(lbi loadBalancerInt) {
	lbi.Test()
}
