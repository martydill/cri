package crypto_api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

type CoinGeckoApi struct {
}

func (api CoinGeckoApi) GetPriceHistory(ticker string, currency string, days int) []HistoricalPrice {
	url := fmt.Sprintf("https://api.coingecko.com/api/v3/coins/%s/market_chart?vs_currency=%s&days=%d", ticker, currency, days)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseData interface{}
	json.Unmarshal(body, &responseData)
	data := responseData.(map[string]interface{})
	historicalData, ok := data["prices"]
	if !ok {
		log.Fatal(data["error"])
	}

	dayData := historicalData.([]interface{})
	var results []HistoricalPrice
	for i, x := range dayData {
		parts := x.([]interface{})

		// API returns hourly data for < 90 days, so let's just
		// convert that to daily data by only including every 24th record
		if days > 90 || i%24 == 0 {
			timestamp := parts[0].(float64) / 1000
			t := time.Unix(int64(timestamp), 0)
			value := parts[1].(float64)
			results = append(results, HistoricalPrice{t, value})
		}
	}
	return results
}

func (api CoinGeckoApi) GetPrice(ticker, currency string) string {
	url := fmt.Sprintf("https://api.coingecko.com/api/v3/simple/price?ids=%s&vs_currencies=%s", ticker, currency)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseData interface{}
	json.Unmarshal(body, &responseData)
	data := responseData.(map[string]interface{})
	tickerData, ok := data[ticker].(map[string]interface{})
	if !ok {
		log.Fatalf("Could not find symbol '%s'", ticker)
	}
	price := tickerData[currency].(float64)
	return strconv.FormatFloat(price, 'f', -1, 64)
}
