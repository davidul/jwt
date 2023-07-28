package cmd

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/spf13/cobra"
	"jwt/pkg"
)

var (
	TokenPath     string
	TokenName     string
	Secret        string
	Claims        map[string]string
	signingMethod string

	genJwtCmd = &cobra.Command{
		Use:   "gen",
		Short: "Generate JWT",
		Long: "Generate JWT token. There are symmetric and asymmetric cipher." +
			"Symmetric cipher uses the same key for encryption and decryption." +
			"Asymmetric cipher uses different key for encryption and decryption." +
			"Symmetric cipher are: HS256, HS384, HS512",
		Run: func(cmd *cobra.Command, args []string) {
			//pathFlag := cmd.Flag("tokenpath")
			secretFlag := cmd.Flag("secret")
			signingMethod := cmd.Flag("signingmethod")
			claimMap, err := cmd.Flags().GetStringToString("claims")

			if err != nil {
				fmt.Println(err)
			}

			if secretFlag != nil && len(secretFlag.Value.String()) > 0 {
				fmt.Printf("=== Generating JWT token with secret === \"%s\" \n", secretFlag.Value.String())

				if signingMethod.Value.String() == "HS256" {
					symmetric, token := pkg.GenerateSymmetric(secretFlag.Value.String(), claimMap, jwt.SigningMethodHS256)
					fmt.Printf("%s \n", pkg.HeaderToString(token))
					fmt.Println(symmetric)
				}

				if signingMethod.Value.String() == "HS384" {
					symmetric, token := pkg.GenerateSymmetric(secretFlag.Value.String(), claimMap, jwt.SigningMethodHS384)
					fmt.Printf("%s \n", pkg.HeaderToString(token))
					fmt.Println(symmetric)
				}

				return
			}

			fmt.Println("=== Generating Simple Token ===")

			if signingMethod.Value.String() == "HS256" {
				signedString, token := pkg.GenerateSimple(claimMap, jwt.SigningMethodHS256)
				fmt.Printf("%s \n", pkg.HeaderToString(token))
				fmt.Printf("Signed string: \n%s\n", signedString)
			}

			if signingMethod.Value.String() == "HS384" {
				signedString, token := pkg.GenerateSimple(claimMap, jwt.SigningMethodHS384)
				fmt.Printf("%s \n", pkg.HeaderToString(token))
				fmt.Printf("Signed string: \n%s\n", signedString)
			}

			if signingMethod.Value.String() == "HS512" {
				signedString, token := pkg.GenerateSimple(claimMap, jwt.SigningMethodHS512)
				fmt.Printf("%s \n", pkg.HeaderToString(token))
				fmt.Printf("Signed string: \n%s\n", signedString)
			}
		},
	}
)

func init() {
	genJwtCmd.Flags().StringVarP(&TokenPath, "tokenpath", "t", "", "Token path")
	genJwtCmd.Flags().StringVarP(&TokenName, "tokenname", "n", "", "Token name")
	genJwtCmd.Flags().StringVarP(&Secret, "secret", "s", "", "hash input key")
	genJwtCmd.Flags().StringToString("claims", nil, "key=value pairs, separated by comma")
	genJwtCmd.Flags().StringVarP(&signingMethod, "signingmethod", "m", "HS256", "signing method")
}
