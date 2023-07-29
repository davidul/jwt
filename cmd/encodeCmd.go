package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"jwt/pkg"
)

var SecretE string

var encodeCmd = &cobra.Command{
	Use:   "encode",
	Short: "Encode JWT token",
	Long:  "Encode JWT token, if secret is not provided, default secret is used",
	Run: func(cmd *cobra.Command, args []string) {
		secret := cmd.Flag("secret")
		strSecret := cmd.Flag("secret").Value.String()
		if strSecret == "" {
			strSecret = pkg.DEFAULT_SECRET
		}

		if len(args) == 0 {
			fmt.Println("Error: No token provided")
			return
		}

		fmt.Printf("=== Encoding JWT token with secret === \"%s\" \n", secret.Value.String())
		encode, err := pkg.Encode(args[0], secret.Value.String())
		if err != nil {
			fmt.Fprintln(cmd.OutOrStderr(), err)
		}
		fmt.Println(encode)
		return
	},
}

func init() {
	encodeCmd.Flags().StringVarP(&SecretE, "secret", "s", "", "secret key")
}
