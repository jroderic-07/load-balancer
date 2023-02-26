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
	backendSize int
	backendCounter int
	mu sync.Mutex
}

func NewRoundRobin(port string, backends *backend.BackendCollection) *RoundRobin {
	roundRobin := &RoundRobin{}
	roundRobin.LoadBalancer.setPort(port)
	roundRobin.LoadBalancer.setBackends(backends.CreateIterator())
	roundRobin.setBackendSize(len(roundRobin.LoadBalancer.backends.Backends))

	return roundRobin
}

func (r *RoundRobin) setBackendSize(backendSize int){
	r.backendSize = backendSize
}

func (r *RoundRobin) checkCounter() {
	if r.backendCounter >= r.backendSize {
		r.backendCounter = 0
		r.LoadBalancer.backends.ResetIndex()
	}
}

func (r *RoundRobin) Serve() {
	r.LoadBalancer.Serve(r)
}

func (r *RoundRobin) lbHandler(res http.ResponseWriter, req *http.Request) {
	r.mu.Lock()
	
	currentBackend := r.LoadBalancer.backends.GetNext()
	for (currentBackend.GetHealth() == false) {
		r.backendCounter++
		r.checkCounter()
		currentBackend = r.LoadBalancer.backends.GetNext()
	}

	currentURL := currentBackend.GetURL()

	r.mu.Unlock()

	r.backendCounter++
	r.checkCounter()

	parsedURL, _ := url.Parse(currentURL)
	reverseProxy := httputil.NewSingleHostReverseProxy(parsedURL)

	reverseProxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, e error){
		currentBackend.SetUnhealthy()
	}

	req.Host = req.URL.Host 

	reverseProxy.ServeHTTP(res, req)
}
