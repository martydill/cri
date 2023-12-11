package cmd

import (
	"fmt"

	"github.com/guptarohit/asciigraph"
	"github.com/spf13/cobra"
)

var days int

var chartCmd = &cobra.Command{
	Use:   "chart",
	Short: "Displays a chart of historical token prices",
	Args:  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),

	Run: func(cmd *cobra.Command, args []string) {
		currency := rootCmd.PersistentFlags().Lookup("currency").Value.String()
		results := api.GetPriceHistory(args[0], currency, days)
		var values []float64
		for i := 0; i < len(results); i++ {
			v := results[i].Price
			values = append(values, v)
		}

		graph := asciigraph.Plot(values, asciigraph.Precision(0), asciigraph.SeriesColors(
			asciigraph.Red,
		), asciigraph.Height(10), asciigraph.Width(80))

		fmt.Println(graph)
	},
}

func init() {
	rootCmd.AddCommand(chartCmd)
	chartCmd.Flags().IntVarP(&days, "days", "d", 90, "Number of days to fetch chart data for")
}
