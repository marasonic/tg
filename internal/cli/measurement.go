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

package cli

import (
	"fmt"
	"io/ioutil"
	"tg/internal/data"
	"tg/internal/http"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

// ... existing measurementCmd definition ...

var sendMeasurementCmd = &cobra.Command{
	Use:   "send",
	Short: "Send measurement data",
	Run: func(cmd *cobra.Command, args []string) {
		entityID, _ := cmd.Flags().GetString("entity-id")
		year, _ := cmd.Flags().GetInt("year")
		value, _ := cmd.Flags().GetFloat64("value")
		random, _ := cmd.Flags().GetBool("random")
		configPath, _ := cmd.Flags().GetString("config")
		backendURL, _ := cmd.Flags().GetString("backend-url")

		var configData map[string]interface{}
		if configPath != "" {
			yamlFile, err := ioutil.ReadFile(configPath)
			if err != nil {
				fmt.Printf("Error reading config file: %v\n", err)
				return
			}
			err = yaml.Unmarshal(yamlFile, &configData)
			if err != nil {
				fmt.Printf("Error unmarshalling config file: %v\n", err)
				return
			}
		}

		days := data.GetDaysInYear(year)
		for _, day := range days {
			var measurementValue float64
			if random {
				measurementValue = data.GenerateRandomValue(1, 100)
			} else {
				measurementValue = value
			}

			payload := make(map[string]interface{})
			for k, v := range configData {
				payload[k] = v
			}
			payload["entity_id"] = entityID
			payload["timestamp"] = day.Format(time.RFC3339)
			payload["value"] = measurementValue

			url := fmt.Sprintf("%s/measurements", backendURL)
			if err := http.SendPostRequest(url, payload); err != nil {
				fmt.Printf("Failed to send measurement for %s: %v\n", day.Format("2006-01-02"), err)
			}
		}
	},
}

// ... existing init function ...


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
