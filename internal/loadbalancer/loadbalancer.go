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
	port int
	backends *backend.BackendCollection
}

func (l *LoadBalancer) setPort(portNo int) {
	l.port = portNo
}

func (l *LoadBalancer) setBackends(backends *backend.BackendCollection) {
	l.backends = backends
}

func (l *LoadBalancer) getPort() int {
	return l.port
}

func (l *LoadBalancer) getBackends() *backend.BackendCollection {
	return l.backends
}

func (l *LoadBalancer) Serve(lbi loadBalancer) {
	s := http.Server {
		Addr: ":8080",
		Handler: http.HandlerFunc(lbi.lbHandler),
	}
	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err.Error())
	}
}
