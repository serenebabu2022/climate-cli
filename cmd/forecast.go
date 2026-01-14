package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/serenebabu2022/climate-cli/internal/api"
)

var (
	lat float64
	lon float64
)

var forecastCmd = &cobra.Command{
	Use:   "forecast",
	Short: "Fetch current weather (Open-Meteo)",
	RunE: func(cmd *cobra.Command, args []string) error {
		resp, err := api.GetForecast(
			"https://api.open-meteo.com/v1/forecast",
			lat,
			lon,
		)
		if err != nil {
			return err
		}

		out, _ := json.MarshalIndent(resp, "", "  ")
		fmt.Println(string(out))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(forecastCmd)
	forecastCmd.Flags().Float64Var(&lat, "lat", 53.3498, "Latitude")
	forecastCmd.Flags().Float64Var(&lon, "lon", -6.2603, "Longitude")
}
