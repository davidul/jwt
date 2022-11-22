package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	TokenPath string
	TokenName string
	Secret    string
	Claims    map[string]string

	genJwtCmd = &cobra.Command{
		Use:   "gen",
		Short: "Generate JWT",
		Run: func(cmd *cobra.Command, args []string) {
			//pathFlag := cmd.Flag("tokenpath")
			secretFlag := cmd.Flag("secret")
			claimMap, err := cmd.Flags().GetStringToString("claims")

			if err != nil {
				fmt.Println(err)
			}

			if secretFlag != nil && len(secretFlag.Value.String()) > 0 {
				fmt.Printf("=== Generating JWT token with secret === \"%s\" \n", secretFlag.Value.String())
				symmetric, token := GenerateSymmetric(secretFlag.Value.String(), claimMap)
				fmt.Printf("%s \n", HeaderToString(token))
				fmt.Println(symmetric)
				return
			}

			fmt.Println("=== Generating Simple Token ===")

			signedString, token := GenerateSimple(claimMap)
			fmt.Printf("%s \n", HeaderToString(token))
			fmt.Printf("Signed string: \n%s\n", signedString)
		},
	}
)

func init() {
	genJwtCmd.Flags().StringVarP(&TokenPath, "tokenpath", "t", "", "Token path")
	genJwtCmd.Flags().StringVarP(&TokenName, "tokenname", "n", "", "Token name")
	genJwtCmd.Flags().StringVarP(&Secret, "secret", "s", "", "Secret")
	genJwtCmd.Flags().StringToString("claims", Claims, "maps")
}
