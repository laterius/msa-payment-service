package service

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) Service {
	return Service{db: db}
}

// StorePayment Реализация методов обращения в базу данных
func (s *Service) StorePayment(orderId uuid.UUID, amount int) error {
	err := s.db.Create(Payment{
		OrderId: orderId,
		Amount:  amount,
	}).Error
	return err
}

func (s *Service) DeletePayment(orderId uuid.UUID) error {
	return s.db.Delete(&Payment{}, orderId).Error
}

type Payment struct {
	OrderId uuid.UUID `json:"orderId" gorm:"type:uuid; not null; unique"`
	Amount  int       `json:"amount"`
}
