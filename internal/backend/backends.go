package backend

import (
	"sync"
)

type Backend struct {
	url string
	healthy bool
	mu sync.RWMutex
}

func New(config_url string) *Backend {
	return &Backend{url: config_url, healthy: true}
}

func (b *Backend) GetURL () string{
	return b.url
}

func (b *Backend) GetHealth() bool {
	return b.healthy
}

func (b *Backend) SetUnhealthy() {
	b.healthy = false
}
