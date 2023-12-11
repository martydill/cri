package cmd

import (
	"cri/crypto_api"
	"fmt"

	"github.com/spf13/cobra"
)

var api = &crypto_api.CoinGeckoApi{}

var priceCmd = &cobra.Command{
	Use:   "price",
	Short: "Fetch the price of a cryptocurrency",
	Args:  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		currency := rootCmd.PersistentFlags().Lookup("currency").Value.String()
		fmt.Println(api.GetPrice(args[0], currency))
	},
}

func init() {
	rootCmd.AddCommand(priceCmd)
}
