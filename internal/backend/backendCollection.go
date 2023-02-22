package backend

import (
	"load_balancer/internal/config"
)

type collection interface {
	createIterator() iterator
}

type backendCollection struct {
	Backends []*Backend
}

func (b *backendCollection) createIterator() iterator {
	return &backendIterator{
		backends: b.Backends,
	}
}

func NewBackendCollection(config *config.Config) *backendCollection {
	var backendSlice []*Backend

	backends := config.GetBackends()

	for i := 0; i < len(backends); i++ {
		backend := New(backends[i], true)
		backendSlice = append(backendSlice, backend)
	}

	return &backendCollection{Backends: backendSlice}
}
