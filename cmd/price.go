package cmd

import (
	"cri/crypto_api"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var api = &crypto_api.CoinGeckoApi{}

var priceCmd = &cobra.Command{
	Use:   "price",
	Short: "Fetch the price of a cryptocurrency",
	Args:  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		currency := rootCmd.PersistentFlags().Lookup("currency").Value.String()
		symbols := strings.Split(args[0], ",")
		results := api.GetPrice(symbols, currency)
		for _, result := range results {
			fmt.Printf("%s: %f\n", result.Coin, result.Price)
		}
	},
}

func init() {
	rootCmd.AddCommand(priceCmd)
}
