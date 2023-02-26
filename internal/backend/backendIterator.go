package backend

type iterator interface {
	hasNext() bool
	GetNext() *Backend
	ResetIndex()
}

type BackendIterator struct {
	index int
	Backends []*Backend
}

func (b *BackendIterator) hasNext() bool {
	if b.index < len(b.Backends){
		return true
	}

	return false
}

func (b *BackendIterator) GetNext() *Backend {
	if b.hasNext() {
		backend := b.Backends[b.index]
		b.index++
		return backend
	}
	
	return nil
}

func (b *BackendIterator) ResetIndex() {
	b.index = 0
}
