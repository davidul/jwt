package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "jwt",
	Short: "Sing JWT",
	Long:  "CLI utility to generate and sing the JWT tokens",
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(genCmd)
	rootCmd.AddCommand(genJwtCmd)
}
