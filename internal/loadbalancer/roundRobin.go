package loadbalancer

type RoundRobin struct {
	LoadBalancer
}

func NewRoundRobin() *RoundRobin {
	return &RoundRobin{}
}

func (r *RoundRobin) Test() {
	print("tes tickles")
}
