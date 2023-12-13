package cmd

import (
	"fmt"
	"strings"

	"github.com/guptarohit/asciigraph"
	"github.com/spf13/cobra"
)

var days int

var chartCmd = &cobra.Command{
	Use:   "chart",
	Short: "Displays a chart of historical token prices",
	Args:  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),

	Run: func(cmd *cobra.Command, args []string) {
		symbols := strings.Split(args[0], ",")
		currency := rootCmd.PersistentFlags().Lookup("currency").Value.String()
		results := api.GetPriceHistory(symbols, currency, days)

		var valuesToPlot [][]float64
		for i := 0; i < len(results); i++ {
			resultsForCoin := results[i]
			var temp []float64
			for j := 0; j < len(resultsForCoin); j++ {
				price := resultsForCoin[j].Price
				temp = append(temp, price)
			}
			valuesToPlot = append(valuesToPlot, temp)
		}
		graph := asciigraph.PlotMany(valuesToPlot, asciigraph.Precision(2), asciigraph.SeriesColors(
			asciigraph.Red,
			asciigraph.Blue,
			asciigraph.Green,
			asciigraph.Yellow,
			asciigraph.Purple,
			asciigraph.Orange,
			asciigraph.Pink,
			asciigraph.Gray,
		), asciigraph.Height(20), asciigraph.Width(80), asciigraph.LegendTexts(symbols...))

		fmt.Println(graph)
	},
}

func init() {
	rootCmd.AddCommand(chartCmd)
	chartCmd.Flags().IntVarP(&days, "days", "d", 90, "Number of days to fetch chart data for")
}
