package cmd

import (
	"github.com/spf13/cobra"
)

var (
	KeyPath string
	KeyName string

	genCmd = &cobra.Command{
		Use:   "genkeys",
		Short: "Generates public/private keypair",
		Run: func(cmd *cobra.Command, args []string) {
			PkRsa(KeyPath, KeyName)
		}}
)

func init() {
	genCmd.Flags().StringVarP(&KeyPath, "keypath", "k", "", "target directory")
	genCmd.Flags().StringVarP(&KeyName, "keyname", "n", "", "certificate name")
	err := genCmd.MarkFlagRequired("keypath")
	if err != nil {
		return
	}
	err = genCmd.MarkFlagRequired("keyname")
	if err != nil {
		return
	}
}
