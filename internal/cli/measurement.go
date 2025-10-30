package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var measurementCmd = &cobra.Command{
	Use:   "measurement",
	Short: "Manage measurements",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use 'tg measurement send' to send measurement data.")
	},
}

var sendMeasurementCmd = &cobra.Command{
	Use:   "send",
	Short: "Send measurement data",
	Run: func(cmd *cobra.Command, args []string) {
		entityID, _ := cmd.Flags().GetString("entity-id")
		year, _ := cmd.Flags().GetInt("year")
		value, _ := cmd.Flags().GetFloat64("value")
		random, _ := cmd.Flags().GetBool("random")
		config, _ := cmd.Flags().GetString("config")
		backendURL, _ := cmd.Flags().GetString("backend-url")

		fmt.Printf("Sending measurement data for entity %s in year %d\n", entityID, year)
		if random {
			fmt.Println("Using random values.")
		} else {
			fmt.Printf("Using fixed value: %f\n", value)
		}
		fmt.Printf("Config file: %s\n", config)
		fmt.Printf("Backend URL: %s\n", backendURL)
		// Data generation and HTTP request logic will be added here.
	},
}

func init() {
	measurementCmd.AddCommand(sendMeasurementCmd)
	sendMeasurementCmd.Flags().String("entity-id", "", "ID of the entity")
	sendMeasurementCmd.Flags().Int("year", 2020, "Year to generate data for")
	sendMeasurementCmd.Flags().Float64("value", 0, "A fixed value to send for each day")
	sendMeasurementCmd.Flags().Bool("random", false, "Use random values between 1-100")
	sendMeasurementCmd.Flags().String("config", "", "Path to a template file for measurement parameters")
	sendMeasurementCmd.Flags().String("backend-url", "http://localhost:8080", "Backend API URL")
	sendMeasurementCmd.MarkFlagRequired("entity-id")
}
