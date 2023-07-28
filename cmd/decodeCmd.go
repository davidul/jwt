package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"jwt/pkg"
)

var decodeCmd = &cobra.Command{
	Use:   "decode token",
	Short: "Decodes JWT token",
	Long:  "Decodes JWT token, if secret is not provided, default secret is used",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Error: No token provided")
			err := cmd.Help()
			if err != nil {
				return
			}

			return
		}

		strSecret := cmd.Flag("secret").Value.String()
		if strSecret == "" {
			strSecret = pkg.DEFAULT_SECRET
		}

		strOutput := cmd.Flag("output").Value.String()
		if strOutput == "" {
			strOutput = "json"
		}
		parse := pkg.Parse(args[0], strSecret)
		// convert jwt.Token to JSON

		if strOutput == "text" {
			fmt.Println(pkg.HeaderToString(parse))
			return
		}

		marshal, _ := json.MarshalIndent(parse.Header, "", "  ")
		fmt.Println(string(marshal))
		indent, _ := json.MarshalIndent(parse.Claims, "", "  ")
		fmt.Println(string(indent))

	},
}

func init() {
	decodeCmd.Flags().String("secret", "", "optional secret key")
	decodeCmd.Flags().String("output", "", "output format json or text")
}
