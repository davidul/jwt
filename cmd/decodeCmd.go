package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/spf13/cobra"
	"jwt/pkg"
)

var decodeCmd = &cobra.Command{
	Use:   "decode token",
	Short: "Decodes JWT token",
	Long: "Decodes JWT token, if secret is not provided, default secret is used" +
		"if public key is provided, token is decoded with public key" +
		"if output is not provided, json is used" +
		"secret or public key is optional, it is used only for validation, not for decoding",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Error: No token provided")
			err := cmd.Help()
			if err != nil {
				return
			}
			return
		}

		//eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJhdWQiLCJleHAiOjE3MjIxNzg5MjksImlhdCI6MTY5MDM4MzcyOSwiaXNzIjoiaXNzIiwibmJmIjoxNjkwNDcwMTI5LCJzdWIiOiJzdWIifQ.puww8DUW_MhVUUzEBUmJf-t7j0jnJlYcF3ftD2BmJLJINZpfnTAwdoeFf7y0n4Hd0nAO7QKql6XN0PqlIRdph8LQr-SR_WXNVUe_8trfmQA-Zxrp-M8WCLV8msgt8waDs6_uXmi1IJiOJVB2ryNs2tEZhwLztifGN1TCU8YU2sbkP9g_Yz7zOw6BFulWiv-am2eHbxMOQeE16-i3in_JpLqT-ypn6o5zNNiYKyVFGeDftKNXk5bQPnDmWg_5mwkZi1ybqGJdy6RsUGQ8PBMPGKsM7JCvrQw8DQEcDMMQ--nZLNtkqk0BHxM7VAG-Vgs7Hz2JFQLmFQKwXgHwRt_ojg

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

		parse := pkg.Parse(args[0], strSecret)
		output(parse, strOutput)

	},
}

func init() {
	decodeCmd.Flags().String("secret", "", "optional secret key")
	decodeCmd.Flags().String("output", "", "output format json or text")
	decodeCmd.Flags().String("publickey", "", "public key file path")
}

func output(token *jwt.Token, outputType string) {
	if outputType == "text" {
		fmt.Println(pkg.HeaderToString(token))
		return
	} else {
		marshal, _ := json.MarshalIndent(token.Header, "", "  ")
		fmt.Println(string(marshal))
		indent, _ := json.MarshalIndent(token.Claims, "", "  ")
		fmt.Println(string(indent))
	}
}
