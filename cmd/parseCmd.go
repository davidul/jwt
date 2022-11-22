package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var parseCmd = &cobra.Command{
	Use:   "parse",
	Short: "parse",
	Long:  "parse",
	Run: func(cmd *cobra.Command, args []string) {
		parse := Parse(args[0])
		fmt.Println(HeaderToString(parse))

	},
}

func init() {
	//parseCmd.Flags().String("token", "", "usage")
}
