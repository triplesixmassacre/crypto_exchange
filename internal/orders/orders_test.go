package orders

import (
	"testing"
	"time"
)

func TestCreateOrder(t *testing.T) {
	orders_list := []struct {
		test_order  Order
		comment     string
		errExcepted bool
	}{
		{
			Order{UserID: 100, BaseAsset: "ETH", QuoteAsset: "USDT", Type: "MARKET", Side: true, Price: 2000.0, Amount: 1.5, FilledAmount: 0.0, Status: "PENDING", CreatedAt: time.Now(), UpdatedAt: time.Now(), Fee: 0.001},
			"Корректный ордер",
			false,
		},
		{
			Order{UserID: 100, BaseAsset: "ETH", QuoteAsset: "USDT", Type: "NON-MARKET", Side: true, Price: 2000.0, Amount: 1.5, FilledAmount: 0.0, Status: "PENDING", CreatedAt: time.Now(), UpdatedAt: time.Now(), Fee: 0.001},
			"Ошибка Type",
			true,
		},
		{
			Order{UserID: 100, BaseAsset: "ETH", QuoteAsset: "USDT", Type: "MARKET", Side: true, Price: 2000.0, Amount: 1.5, FilledAmount: 0.0, Status: "PENDING", CreatedAt: time.Now(), UpdatedAt: time.Now(), Fee: 0.001},
			"Ошибка Side",
			true,
		},
		{
			Order{UserID: 100, BaseAsset: "ETH", QuoteAsset: "USDT", Type: "MARKET", Side: true, Price: 2000.0, Amount: 1.5, FilledAmount: 0.0, Status: "IDK", CreatedAt: time.Now(), UpdatedAt: time.Now(), Fee: 0.001},
			"Ошибка Status",
			true,
		},
		{
			Order{UserID: 100, BaseAsset: "ETH", QuoteAsset: "USDT", Type: "MARKET", Side: true, Price: -1.0, Amount: 1.5, FilledAmount: 0.0, Status: "PENDING", CreatedAt: time.Now(), UpdatedAt: time.Now(), Fee: 0.001},
			"Price < 0",
			true,
		},
		{
			Order{UserID: 100, BaseAsset: "ETH", QuoteAsset: "USDT", Type: "MARKET", Side: true, Price: 2000.0, Amount: 1.5, FilledAmount: 2, Status: "PENDING", CreatedAt: time.Now(), UpdatedAt: time.Now(), Fee: 0.001},
			"FilledAmount > Amount",
			true,
		},
		{
			Order{UserID: 100, BaseAsset: "ETH", QuoteAsset: "USDT", Type: "MARKET", Side: true, Price: 2000.0, Amount: 1.5, FilledAmount: -1, Status: "PENDING", CreatedAt: time.Now(), UpdatedAt: time.Now(), Fee: 0.001},
			"FilledAmount < 0",
			true,
		},
	}

	for _, tt := range orders_list {
		err := tt.test_order.Validate()
		if (err != nil) != tt.errExcepted {
			t.Errorf("ошибка в %v: %s", tt.comment, err)
		}
	}
}
