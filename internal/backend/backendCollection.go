package backend

import (
	"github.com/jroderic-07/load-balancer/internal/config"
)

type collection interface {
	createIterator() iterator
}

type BackendCollection struct {
	Backends []*Backend
}

func (b *BackendCollection) createIterator() iterator {
	return &backendIterator{
		backends: b.Backends,
	}
}

func NewBackendCollection(config *config.Config) *BackendCollection {
	var backendSlice []*Backend

	backends := config.GetBackends()

	for i := 0; i < len(backends); i++ {
		backend := New(backends[i], true)
		backendSlice = append(backendSlice, backend)
	}

	return &BackendCollection{Backends: backendSlice}
}

func (b *BackendCollection) GetBackends() []*Backend {
	return b.Backends
}
