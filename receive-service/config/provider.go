package config

import "receive-service/worker"

type IProvider interface {
	Initialize()
}

type Provider struct {
	worker worker.IWorker
}

func NewProvider() IProvider {
	return &Provider{
		worker: worker.NewWorker(),
	}
}

func (p *Provider) Initialize() {
	p.worker.StartWorker()
}
