package loadbalancer

import (
	"net/http"
	"log"

	"github.com/jroderic-07/load-balancer/internal/backend"	
)

type loadBalancer interface {
	lbHandler(w http.ResponseWriter, rq *http.Request)
	Serve()
}

type LoadBalancer struct {
	port string
	backends *backend.BackendIterator
}

func (l *LoadBalancer) setPort(portNo string) {
	l.port = portNo
}

func (l *LoadBalancer) setBackends(backends *backend.BackendIterator) {
	l.backends = backends
}

func (l *LoadBalancer) Serve(lbi loadBalancer) {
	s := http.Server {
		Addr: l.port, 
		Handler: http.HandlerFunc(lbi.lbHandler),
	}
	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err.Error())
	}
}
