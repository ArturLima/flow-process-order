package config

import "order-status-log-service/worker"

type IProvider interface {
	Initialize()
}

type Provider struct {
	worker worker.IWorker
}

func NewWorker() IProvider {
	return &Provider{
		worker: worker.NewWorker(),
	}
}

func (p *Provider) Initialize() {
	p.worker.StartWorker()
}
