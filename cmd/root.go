package cmd

import (
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var dateFormat string
var rootCmd = &cobra.Command{
	Use:   "jwt",
	Short: "Sign JWT",
	Long:  "CLI utility to generate and sign the JWT tokens",
}

var Logger *zap.Logger

func Execute(logger *zap.Logger) error {
	Logger = logger
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&dateFormat, "dateformat", "d", "RFC",
		"pass datetime format")
	rootCmd.AddCommand(genPkCmd)
	rootCmd.AddCommand(genJwtCmd)
	rootCmd.AddCommand(decodeCmd)
	rootCmd.AddCommand(encodeCmd)
}
