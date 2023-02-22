package loadbalancer

import (
	"load_balancer/internal/backend"
)

type LoadBalancer struct {
	port int
	backends []*backend.Backend
}

func New(port int, backend []*backend.Backend) *LoadBalancer {
	return &LoadBalancer{port, backend}
}

func (l *LoadBalancer) GetPort() int {
	return l.port
}

func (l *LoadBalancer) GetBackends() []*backend.Backend {
	return l.backends
}
