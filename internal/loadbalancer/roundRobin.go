package loadbalancer

import (
	"sync"
	"net/http/httputil"
	"net/http"
	"net/url"

	"github.com/jroderic-07/load-balancer/internal/backend"
)

type RoundRobin struct {
	LoadBalancer
	backendCounter int
	mu sync.Mutex
}

func NewRoundRobin(port string, backends *backend.BackendCollection) *RoundRobin {
	roundRobin := &RoundRobin{backendCounter: 0}
	roundRobin.LoadBalancer.setPort(port)
	roundRobin.LoadBalancer.setBackends(backends)

	return roundRobin
}

func (r *RoundRobin) lbHandler(res http.ResponseWriter, req *http.Request) {
	backendsLength := len(r.LoadBalancer.backends.Backends)
	
	r.mu.Lock()

	currentURL := r.LoadBalancer.backends.Backends[r.backendCounter].GetURL()

	r.mu.Unlock()

	r.backendCounter++
	if r.backendCounter >= backendsLength {
		r.backendCounter = 0
	}
	
	parsedURL, _ := url.Parse(currentURL)
	reverseProxy := httputil.NewSingleHostReverseProxy(parsedURL)

	req.Host = req.URL.Host 

	reverseProxy.ServeHTTP(res, req)
}
