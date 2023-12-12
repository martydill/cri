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

		var values [][]float64

		for i := 0; i < len(results); i++ {
			v := results[i]
			var temp []float64
			for j := 0; j < len(v); j++ {
				temp = append(temp, v[j].Price)
			}
			values = append(values, temp)
		}
		graph := asciigraph.PlotMany(values, asciigraph.Precision(2), asciigraph.SeriesColors(
			asciigraph.Red,
			asciigraph.Blue,
		), asciigraph.Height(20), asciigraph.Width(80), asciigraph.LowerBound(0))

		fmt.Println(graph)
	},
}

func init() {
	rootCmd.AddCommand(chartCmd)
	chartCmd.Flags().IntVarP(&days, "days", "d", 90, "Number of days to fetch chart data for")
}
