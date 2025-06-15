package orders

import (
	"fmt"
	"time"
)

type OrderType string

const (
	OrderTypeLimit  OrderType = "LIMIT"
	OrderTypeMarket OrderType = "MARKET"
)

func (ot OrderType) IsValid() bool {
	switch ot {
	case OrderTypeLimit, OrderTypeMarket:
		return true
	default:
		return false
	}
}

type OrderStatus string

const (
	OrderStatusPending  OrderStatus = "PENDING"
	OrderStatusPatrial  OrderStatus = "PARTIALLY_FILLED"
	OrderStatusFilled   OrderStatus = "FILLED"
	OrderStatusCanceled OrderStatus = "CANCELED"
)

func (ost OrderStatus) IsValid() bool {
	switch ost {
	case OrderStatusPending, OrderStatusPatrial, OrderStatusFilled, OrderStatusCanceled:
		return true
	default:
		return false
	}
}

type Order struct {
	ID           int64
	UserID       int64
	BaseAsset    string
	QuoteAsset   string
	Type         OrderType
	Side         bool // 1 - покупка, 0 - продажа
	Price        float64
	Amount       float64
	FilledAmount float64
	Status       OrderStatus
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Fee          float64 // Значение от 0 до 1
}

func (o *Order) Validate() error {
	if !o.Type.IsValid() {
		return fmt.Errorf("неккоректное указание Type: %v", o.Type)
	}
	if !o.Status.IsValid() {
		return fmt.Errorf("неккоректное указание Status: %v", o.Status)
	}
	if o.Price < 0 {
		return fmt.Errorf("цена должна быть больше 0: %v", o.Price)
	}
	if o.Amount <= 0 {
		return fmt.Errorf("количество должно быть больше 0: %v", o.Amount)
	}
	if o.FilledAmount < 0 || o.FilledAmount > o.Amount {
		return fmt.Errorf("не может быть частичное выполнение ордера меньше 0 или больше количества токенов: %v", o.FilledAmount)
	}
	return nil
}
