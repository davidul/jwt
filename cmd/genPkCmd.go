package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"jwt/pkg"
)

var (
	KeyPath        string
	PrivateKeyName string
	PublicKeyName  string
	KeyType        string

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
			privateKeyName := cmd.Flag("privatekey")
			publicKeyName := cmd.Flag("publickey")

			output := cmd.Flag("output")
			if output.Value.String() != "" {
				strOutput = output.Value.String()
			}

			if keyPath.Value.String() != "" {
				strOutput = "file"
			}

			if privateKeyName.Value.String() != "" {
				strPrivateKeyName = privateKeyName.Value.String()
			}

			if publicKeyName.Value.String() != "" {
				strPublicKeyName = publicKeyName.Value.String()
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
					pkg.EncodePrivateKeyToPemFile(bPrivateKey, keyPath.Value.String(), strPrivateKeyName)
					pkg.EncodePublicKeyToPemFile(bPublicKey, keyPath.Value.String(), strPublicKeyName)
				}
			} else if strKeyTypeName == "ecdsa" {
				privateKey, publicKey := pkg.GenKeysEcdsa()
				bPrivateKey, bPublicKey := pkg.MarshalEcdsa(privateKey, publicKey)
				if strOutput == "stdout" {
					pr, pu := pkg.EncodePem(bPrivateKey, bPublicKey)
					fmt.Println(string(pr))
					fmt.Println(string(pu))
				} else {
					pkg.EncodePemToFile(bPrivateKey, bPublicKey, keyPath.Value.String(), strPrivateKeyName)
				}
			}
		}}
)

func init() {
	genPkCmd.Flags().StringVarP(&KeyPath, "keypath", "k", "", "target directory")
	genPkCmd.Flags().StringVarP(&PrivateKeyName, "privatekey", "n", "", "file name")
	genPkCmd.Flags().StringVarP(&PublicKeyName, "publickey", "p", "", "file name")
	genPkCmd.Flags().StringVarP(&KeyType, "keytype", "t", "rsa", "key type")
	genPkCmd.Flags().StringVarP(&Output, "output", "o", "", "output file or stdout")
}
