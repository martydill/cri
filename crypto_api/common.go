package crypto_api

import (
	"time"
)

type HistoricalPrice struct {
	Date  time.Time
	Price float64
	Coin  string
}

type Price struct {
	Coin  string
	Price float64
}

type ICryptoApi interface {
	GetPrice([]string) []Price
	GetPriceHistory([]string, string, int) [][]HistoricalPrice
}
