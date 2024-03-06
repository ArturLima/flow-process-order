package worker

type IWorker interface {
	StartWorker()
}

type Worker struct {
}

func NewWorker() IWorker {
	return &Worker{}
}
func (w *Worker) StartWorker() {
	println("teste")
	return
}
