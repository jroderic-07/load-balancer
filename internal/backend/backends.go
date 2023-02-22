package backend

type Backend struct {
	url string
	health bool
}

func New(config_url string, config_health bool) *Backend {
	return &Backend{url: config_url, health: config_health}
}

func (b *Backend) GetURL () string{
	return b.url
}

func (b *Backend) GetHealth() bool {
	return b.health
}

func (b *Backend) CheckHealth() {}
