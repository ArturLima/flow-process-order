package services

import (
	"order-status-log-service/core/models"
	"order-status-log-service/infra/db"

	"github.com/google/uuid"
)

type IOrderStatusService interface {
	Insert(os *models.OrderStatus) (err error)
	Get(orderId uuid.UUID) (orderStatus models.OrderStatus, err error)
}

type OrderStatusService struct {
	repo db.IOrderStatusRepository
}

func NewOrderStatusService() IOrderStatusService {
	return &OrderStatusService{
		repo: db.NewOrderRepository(),
	}
}

func (s *OrderStatusService) Insert(os *models.OrderStatus) (err error) {
	err = s.repo.Insert(os)
	return
}

func (s *OrderStatusService) Get(orderId uuid.UUID) (orderStatus models.OrderStatus, err error) {
	os, err := s.repo.Get(orderId)
	return os, err
}
