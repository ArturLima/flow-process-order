package handlers

type IOrderStatusHandler interface {
	Execute(b []byte) (done bool)
}

type OrderStatusHandler struct {
}

func NewOrderStatusHandler() IOrderStatusHandler {
	return &OrderStatusHandler{}
}

func (os *OrderStatusHandler) Execute(b []byte) (done bool) {
	return true
}
