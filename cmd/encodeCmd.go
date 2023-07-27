package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"jwt/pkg"
)

var SecretE string

var encodeCmd = &cobra.Command{
	Use:   "encode",
	Short: "Encode JWT",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Encode JWT")
		secret := cmd.Flag("secret")
		if len(args) == 0 {
			fmt.Println("Error: No token provided")
			return
		}

		if secret != nil && len(secret.Value.String()) > 0 {
			fmt.Printf("=== Encoding JWT token with secret === \"%s\" \n", secret.Value.String())
			encode := pkg.Encode(args[0], secret.Value.String())
			fmt.Println(encode)
			return
		}
	},
}

func init() {
	encodeCmd.Flags().StringVarP(&SecretE, "secret", "s", "", "secret key")
}
