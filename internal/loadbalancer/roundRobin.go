package loadbalancer

import (
	"fmt"
	"errors"
	"sync"
	"net/http/httputil"
	//"log"
	"net/http"
	//"net/http/httputil"
	"net/url"
	"github.com/jroderic-07/load-balancer/internal/backend"
)

type RoundRobin struct {
	LoadBalancer
	backendCounter int
	mu sync.Mutex
}

func NewRoundRobin(port int, backends *backend.BackendCollection) *RoundRobin {
	roundRobin := &RoundRobin{backendCounter: 0}
	roundRobin.LoadBalancer.setPort(port)
	roundRobin.LoadBalancer.setBackends(backends)

	return roundRobin
}

func ModifyResponse(res *http.Response) error {
    if res.StatusCode == 404 {
        return errors.New("404 error from the host")
    }
    return nil
}  

func ErrHandle(res http.ResponseWriter, req *http.Request, err error) {
    fmt.Println(err)
}  

func (r *RoundRobin) lbHandler(res http.ResponseWriter, req *http.Request) {
	backendsLength := len(r.LoadBalancer.backends.Backends)
	
	r.mu.Lock()

	currentURL := r.LoadBalancer.backends.Backends[0].GetURL()
	print(currentURL)

	r.mu.Unlock()

	r.backendCounter++
	if r.backendCounter >= backendsLength {
		r.backendCounter = 0
	}

	parsedURL, _ := url.Parse("http://www.google.com/")
	reverseProxy := httputil.NewSingleHostReverseProxy(parsedURL)

	reverseProxy.ErrorHandler = ErrHandle
    	reverseProxy.ModifyResponse = ModifyResponse

	req.Host = req.URL.Host 

	fmt.Println(parsedURL)

	reverseProxy.ServeHTTP(res, req)
	fmt.Println(*reverseProxy)
}
