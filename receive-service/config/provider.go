package config

type IProvider interface {
	Initialize()
}

type Provider struct {
	worker worker.IWorker
}

func NewProvider() IProvider {
	return &Provider{}
}

func (p *Provider) Initialize() {

}
