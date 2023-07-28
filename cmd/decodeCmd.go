package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"jwt/pkg"
)

var decodeCmd = &cobra.Command{
	Use:   "decode",
	Short: "decode",
	Long:  "decode",
	Run: func(cmd *cobra.Command, args []string) {
		parse := pkg.Parse(args[0], args[1])
		// convert jwt.Token to JSON
		marshal, _ := json.MarshalIndent(parse.Header, "", "  ")
		fmt.Println(string(marshal))
		indent, _ := json.MarshalIndent(parse.Claims, "", "  ")
		fmt.Println(string(indent))
		fmt.Println(pkg.HeaderToString(parse))

	},
}

func init() {
	//decodeCmd.Flags().String("token", "", "usage")
}
