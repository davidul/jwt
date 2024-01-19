package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/spf13/cobra"
	"jwt/pkg"
)

const ErrorNoToken = "Error: No token provided\n"

var decodeCmd = &cobra.Command{
	Use:   "decode token",
	Short: "Decodes JWT token",
	Long: "Decodes JWT token, if secret is not provided, default secret is used" +
		"if public key is provided, token is decoded with public key" +
		"if output is not provided, json is used" +
		"secret or public key is optional, it is used only for validation, not for decoding",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			_, err := fmt.Fprint(cmd.OutOrStderr(), ErrorNoToken)
			if err != nil {
				return
			}

			err = cmd.Help()
			if err != nil {
				return
			}
			return
		}

		strOutput := cmd.Flag("output").Value.String()
		if strOutput == "" {
			strOutput = "json"
		}

		strPublicKey := cmd.Flag("publickey").Value.String()
		if strPublicKey != "" {
			token := pkg.ParseWithPublicKeyFile(args[0], strPublicKey)
			output(token, strOutput)
			return
		}
		strSecret := cmd.Flag("secret").Value.String()
		if strSecret == "" {
			strSecret = pkg.DEFAULT_SECRET
		}

		parse, err := pkg.Parse(args[0], strSecret)
		if err != nil {
			fmt.Fprintln(cmd.OutOrStderr(), err)
			return
		}
		fmt.Fprintln(cmd.OutOrStderr(), output(parse, strOutput))

	},
}

func init() {
	decodeCmd.Flags().String("secret", "", "optional secret key")
	decodeCmd.Flags().String("output", "", "output format json or text")
	decodeCmd.Flags().String("publickey", "", "public key file path")
}

// output token to stdout
// if outputType is text, output only header
// if outputType is json, output header and claims
func output(token *jwt.Token, outputType string) string {
	if outputType == "text" {
		return pkg.HeaderToString(token)
	} else {
		marshal, _ := json.MarshalIndent(token.Header, "", "  ")
		indent, _ := json.MarshalIndent(token.Claims, "", "  ")
		return string(marshal) + string(indent)
	}
}
