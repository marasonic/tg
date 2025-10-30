package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var entityCmd = &cobra.Command{
	Use:   "entity",
	Short: "Manage entities",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use 'tg entity create' to create a new entity.")
	},
}

package cli

import (
	"fmt"
	"tg/internal/http"

	"github.com/spf13/cobra"
)

// ... existing entityCmd definition ...

var createEntityCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new entity",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		entityType, _ := cmd.Flags().GetString("type")
		backendURL, _ := cmd.Flags().GetString("backend-url")

		payload := map[string]string{
			"name": name,
			"type": entityType,
		}

		url := fmt.Sprintf("%s/entities", backendURL)
		if err := http.SendPostRequest(url, payload); err != nil {
			fmt.Printf("Failed to create entity: %v\n", err)
		}
	},
}

func init() {
	entityCmd.AddCommand(createEntityCmd)
	createEntityCmd.Flags().String("name", "", "Name of the entity")
	createEntityCmd.Flags().String("type", "", "Type of the entity")
	createEntityCmd.Flags().String("backend-url", "http://localhost:8080", "Backend API URL")
	createEntityCmd.MarkFlagRequired("name")
}

