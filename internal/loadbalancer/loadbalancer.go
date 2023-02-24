package loadbalancer

import (
	"net/http"
	"log"

	"github.com/jroderic-07/load-balancer/internal/backend"	
)

type loadBalancer interface {
	lbHandler(w http.ResponseWriter, rq *http.Request)
}

type LoadBalancer struct {
	port string
	backends *backend.BackendCollection
}

func (l *LoadBalancer) setPort(portNo string) {
	l.port = portNo
}

func (l *LoadBalancer) setBackends(backends *backend.BackendCollection) {
	l.backends = backends
}

func (l *LoadBalancer) getPort() string {
	return l.port
}

func (l *LoadBalancer) getBackends() *backend.BackendCollection {
	return l.backends
}

func (l *loadBalancer) checkBackends() {}

func (l *LoadBalancer) Serve(lbi loadBalancer) {
	s := http.Server {
		Addr: l.port, 
		Handler: http.HandlerFunc(lbi.lbHandler),
	}
	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err.Error())
	}
}
