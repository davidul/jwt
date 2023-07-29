package cmd

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/spf13/cobra"
	"jwt/pkg"
)

var (
	Secret        string
	signingMethod string

	PrivateKey string

	genJwtCmd = &cobra.Command{
		Use:   "gen",
		Short: "Generate JWT",
		Long: "Generate JWT token. There are symmetric and asymmetric cipher." +
			"Symmetric cipher uses the same key for encryption and decryption." +
			"Asymmetric cipher uses different key for encryption and decryption." +
			"Symmetric cipher are: HS256, HS384, HS512",
		Run: func(cmd *cobra.Command, args []string) {
			secretFlag := cmd.Flag("secret")
			signingMethod := cmd.Flag("signingmethod")
			claimMap, err := cmd.Flags().GetStringToString("claims")
			privateKeyString := cmd.Flag("privatekey").Value.String()

			if privateKeyString != "" {
				fmt.Println("=== Generating JWT token with private key ===")
				genAsymmetric(privateKeyString, claimMap)
				return
			}
			if err != nil {
				fmt.Println(err)
			}

			if secretFlag != nil && len(secretFlag.Value.String()) > 0 {
				fmt.Printf("=== Generating JWT token with secret === \"%s\" \n", secretFlag.Value.String())
				genSymmetric(signingMethod.Value.String(), secretFlag.Value.String(), claimMap)
				return
			}

			fmt.Println("=== Generating Simple Token ===")
			genSimple(signingMethod.Value.String(), claimMap)
		},
	}
)

func init() {
	genJwtCmd.Flags().StringVarP(&Secret, "secret", "s", "", "hash input key")
	genJwtCmd.Flags().StringToString("claims", nil, "key=value pairs, separated by comma")
	genJwtCmd.Flags().StringVarP(&signingMethod, "signingmethod", "m", "HS256",
		"signing method HS256 | HS384 | HS512")
	genJwtCmd.Flags().StringVarP(&PrivateKey, "privatekey", "p", "", "private key")
}

func genSimple(smethod string, claimMap map[string]string) {
	var signedString string
	var token *jwt.Token

	switch smethod {
	case "HS256":
		signedString, token = pkg.GenerateSimple(claimMap, jwt.SigningMethodHS256)
	case "HS384":
		signedString, token = pkg.GenerateSimple(claimMap, jwt.SigningMethodHS384)
	case "HS512":
		signedString, token = pkg.GenerateSimple(claimMap, jwt.SigningMethodHS512)
	default:
		fmt.Println("Error: Invalid signing method")
		fmt.Println("Valid signing methods are: HS256, HS384, HS512")
		return
	}

	fmt.Printf("%s \n", pkg.HeaderToString(token))
	fmt.Printf("Signed string: \n%s\n", signedString)
}

func genSymmetric(smethod string, secret string, claimMap map[string]string) {
	var signedString string
	var token *jwt.Token

	switch smethod {
	case "HS256":
		signedString, token = pkg.GenerateSymmetric(secret, claimMap, jwt.SigningMethodHS256)
	case "HS384":
		signedString, token = pkg.GenerateSymmetric(secret, claimMap, jwt.SigningMethodHS384)
	case "HS512":
		signedString, token = pkg.GenerateSymmetric(secret, claimMap, jwt.SigningMethodHS512)
	default:
		fmt.Println("Error: Invalid signing method")
		fmt.Println("Valid signing methods are: HS256, HS384, HS512")
		return
	}

	fmt.Printf("%s \n", pkg.HeaderToString(token))
	fmt.Printf("Signed string: \n%s\n", signedString)
}

func genAsymmetric(privateKeyString string, claimMap map[string]string) {
	var signedString string
	//var token *jwt.Token

	pemBlock := pkg.DecodePrivatePemFromFile(privateKeyString)
	rsa := pkg.UnmarshalPrivateRsa(pemBlock)
	signedString = pkg.GenerateSigned(claimMap, rsa)

	//fmt.Printf("%s \n", pkg.HeaderToString(token))
	fmt.Printf("Signed string: \n%s\n", signedString)
}
