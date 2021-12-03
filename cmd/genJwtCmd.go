package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	TokenPath string
	TokenName string
	Secret    string

	genJwtCmd = &cobra.Command{
		Use:   "genjwt",
		Short: "Generate JWT",
		Run: func(cmd *cobra.Command, args []string) {
			//pathFlag := cmd.Flag("tokenpath")
			secretFlag := cmd.Flag("secret")

			if secretFlag != nil {
				fmt.Println(GenerateSymmetric(secretFlag.Value.String()))
			}
		},
	}
)

func init() {
	genJwtCmd.Flags().StringVarP(&TokenPath, "tokenpath", "t", "", "Token path")
	genJwtCmd.Flags().StringVarP(&TokenName, "tokenname", "n", "", "Token name")
	genJwtCmd.Flags().StringVarP(&Secret, "secret", "s", "", "Secret")
}
