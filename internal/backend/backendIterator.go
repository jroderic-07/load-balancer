package backend

type iterator interface {
	hasNext() bool
	getNext() *Backend
}

type backendIterator struct {
	index int
	backends []*Backend
}

func (b *backendIterator) hasNext() bool {
	if b.index < len(b.backends){
		return true
	}

	return false
}

func (b *backendIterator) getNext() *Backend {
	if b.hasNext() {
		backend := b.backends[b.index]
		b.index++
		return backend
	}
	
	return nil
}
