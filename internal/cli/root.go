package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tg",
	Short: "A CLI tool to generate REST requests",
	Long:  `tg is a flexible CLI for sending REST requests to a backend, with support for data generation and traffic scenarios.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to tg! Use 'tg help' to see available commands.")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(entityCmd)
	rootCmd.AddCommand(measurementCmd)
	rootCmd.AddCommand(scenarioCmd)
}

