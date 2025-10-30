package cli

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"text/template"
	"tg/internal/auth"
	"tg/internal/http"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

var scenarioCmd = &cobra.Command{
	Use:   "scenario",
	Short: "Manage scenarios",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use 'tg scenario run' to execute a scenario.")
	},
}

var runScenarioCmd = &cobra.Command{
	Use:   "run",
	Short: "Run a traffic scenario from a file",
	Run: func(cmd *cobra.Command, args []string) {
		filePath, _ := cmd.Flags().GetString("file")
		if filePath == "" {
			fmt.Println("Please provide a scenario file with --file")
			return
		}

		yamlFile, err := ioutil.ReadFile(filePath)
		if err != nil {
			fmt.Printf("Error reading scenario file: %v\n", err)
			return
		}

		var scenario struct {
			Name      string
			Variables map[string]interface{}
			Steps     []struct {
				Name     string
				Request  struct {
					Method string
					URL    string
					Body   map[string]interface{}
				}
				Register string
			}
		}

		err = yaml.Unmarshal(yamlFile, &scenario)
		if err != nil {
			fmt.Printf("Error unmarshalling scenario file: %v\n", err)
			return
		}

		fmt.Printf("Running scenario: %s\n", scenario.Name)

		registeredVars := make(map[string]interface{})

		for _, step := range scenario.Steps {
			fmt.Printf("Executing step: %s\n", step.Name)

			// Substitute variables
			tmpl, err := template.New("url").Parse(step.Request.URL)
			if err != nil {
				fmt.Printf("Error parsing URL template: %v\n", err)
				continue
			}
			var urlBuf bytes.Buffer
			err = tmpl.Execute(&urlBuf, scenario.Variables)
			if err != nil {
				fmt.Printf("Error executing URL template: %v\n", err)
				continue
			}
			url := urlBuf.String()

			// This is a simplified substitution for the body. A real implementation
			// would need to recursively parse the body structure.
			body := step.Request.Body
			for key, val := range body {
				if strVal, ok := val.(string); ok {
					tmpl, err := template.New("body").Parse(strVal)
					if err != nil {
						fmt.Printf("Error parsing body template: %v\n", err)
						continue
					}
					var bodyBuf bytes.Buffer
					err = tmpl.Execute(&bodyBuf, registeredVars)
					if err != nil {
						fmt.Printf("Error executing body template: %v\n", err)
						continue
					}
					body[key] = bodyBuf.String()
				}
			}

			token, err := auth.GetToken()
			if err != nil {
				fmt.Printf("Failed to get token: %v\n", err)
				continue
			}
			if err := http.SendPostRequest(url, token, body); err != nil {
				fmt.Printf("Failed to execute step '%s': %v\n", step.Name, err)
			}
		}
	},
}

func init() {
	scenarioCmd.AddCommand(runScenarioCmd)
	runScenarioCmd.Flags().String("file", "", "Path to the scenario file")
	rootCmd.AddCommand(scenarioCmd)
}
