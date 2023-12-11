package crypto_api

import (
	"time"
)

type HistoricalPrice struct {
	Date  time.Time
	Price float64
}

type ICryptoApi interface {
	GetPrice(string) string
	GetPriceHistory(string, string, int) []HistoricalPrice
}
