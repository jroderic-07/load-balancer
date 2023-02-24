package loadbalancer

type RoundRobin struct {
	LoadBalancer
}

func NewRoundRobin() *RoundRobin {
	return &RoundRobin{}
}

func (r *RoundRobin) lbHandler() {
	print("test")
}
