package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"jwt/pkg"
)

var (
	KeyPath string
	KeyName string
	KeyType string

	Output string

	genPkCmd = &cobra.Command{
		Use:   "genkeys",
		Short: "Generates public/private keypair",
		Long: "Generates public/private keypair.\n" +
			"Default output is stdout.\n" +
			"Key path is a directory where keys will be stored. Key name is a file name without extension.\n" +
			"Keys will be stored in PEM format.\n" +
			"If you omit the keypath, keys will be stored in the current directory.\n" +
			"If you omit the prefix, keys will be stored as private.pem and public.pem.\n",
		Example: "jwt genkeys --keypath . --prefix test",
		Run: func(cmd *cobra.Command, args []string) {
			strPrivateKeyName := "private.pem"
			strPublicKeyName := "public.pem"
			strKeyTypeName := "rsa"
			strOutput := "stdout"

			keyPath := cmd.Flag("keypath")
			keyName := cmd.Flag("prefix")

			output := cmd.Flag("output")
			if output.Value.String() != "" {
				strOutput = output.Value.String()
			}
			if keyName.Value.String() != "" {
				strPrivateKeyName = keyName.Value.String()
				strPrivateKeyName += "_private.pem"

				strPublicKeyName = keyName.Value.String()
				strPublicKeyName += "_public.pem"
			}
			keyType := cmd.Flag("keytype")
			if keyType.Value.String() != "" {
				strKeyTypeName = keyType.Value.String()
			}
			if strKeyTypeName == "rsa" {
				privateKey, publicKey := pkg.GenKeysRsa()
				bPrivateKey, bPublicKey := pkg.MarshalRsa(privateKey, publicKey)
				if strOutput == "stdout" {
					pr, pu := pkg.EncodePem(bPrivateKey, bPublicKey)
					fmt.Println(string(pr))
					fmt.Println(string(pu))
				} else {
					pkg.EncodePemToFile(bPrivateKey, bPublicKey, keyPath.Value.String(), strPrivateKeyName)
				}

			}
			//PkRsa(KeyPath, KeyName)
		}}
)

func init() {
	genPkCmd.Flags().StringVarP(&KeyPath, "keypath", "k", "", "target directory")
	genPkCmd.Flags().StringVarP(&KeyName, "prefix", "n", "", "file prefix")
	genPkCmd.Flags().StringVarP(&KeyType, "keytype", "t", "rsa", "key type")
	genPkCmd.Flags().StringVarP(&Output, "output", "o", "", "output file or stdout")
}
