package crypto_api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

type CoinGeckoApi struct {
}

func (api CoinGeckoApi) GetPriceHistory(coins []string, currency string, days int) [][]HistoricalPrice {
	var results [][]HistoricalPrice

	for _, coin := range coins {
		url := fmt.Sprintf("https://api.coingecko.com/api/v3/coins/%s/market_chart?vs_currency=%s&days=%d", coin, currency, days)
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
			log.Fatalf("Could not fetch price data: %s\n", data["error"])
		}

		dayData := historicalData.([]interface{})
		var resultsForCoin []HistoricalPrice
		for i, x := range dayData {
			parts := x.([]interface{})

			// API returns hourly data for < 90 days, so let's just
			// convert that to daily data by only including every 24th record
			if days > 90 || i%24 == 0 {
				timestamp := parts[0].(float64) / 1000
				t := time.Unix(int64(timestamp), 0)
				value := parts[1].(float64)
				resultsForCoin = append(resultsForCoin, HistoricalPrice{t, value, coin})
			}
		}

		results = append(results, resultsForCoin)
	}
	return results
}

func (api CoinGeckoApi) GetPrice(ticker []string, currency string) []Price {
	url := fmt.Sprintf("https://api.coingecko.com/api/v3/simple/price?ids=%s&vs_currencies=%s", strings.Join(ticker, ","), currency)
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
	var results []Price
	for _, coin := range ticker {
		tickerData, ok := data[coin].(map[string]interface{})
		if ok {
			price, ok := tickerData[currency].(float64)
			if ok {
				results = append(results, Price{Coin: coin, Price: price})
			} else {
				fmt.Println("No data for currency", currency, "for", coin)
			}
		} else {
			fmt.Println("Could not find coin", coin)
		}
	}
	return results
}
