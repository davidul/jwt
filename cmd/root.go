package cmd

import (
	"github.com/spf13/cobra"
)

var dateFormat string
var rootCmd = &cobra.Command{
	Use:   "jwt",
	Short: "Sign JWT",
	Long:  "CLI utility to generate and sign the JWT tokens",
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&dateFormat, "dateformat", "d", "RFC",
		"pass datetime format")
	rootCmd.AddCommand(genCmd)
	rootCmd.AddCommand(genJwtCmd)
	rootCmd.AddCommand(parseCmd)
}
