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

var createEntityCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new entity",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		entityType, _ := cmd.Flags().GetString("type")
		fmt.Printf("Creating entity with name: %s and type: %s\n", name, entityType)
		// HTTP request logic will be added here in a later step.
	},
}

func init() {
	entityCmd.AddCommand(createEntityCmd)
	createEntityCmd.Flags().String("name", "", "Name of the entity")
	createEntityCmd.Flags().String("type", "", "Type of the entity")
	createEntityCmd.MarkFlagRequired("name")
}
