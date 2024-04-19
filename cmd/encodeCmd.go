package cmd

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"io/fs"
	"jwt/pkg"
	"os"
)

var SecretE string

var file string

var key string

var encodeCmd = &cobra.Command{
	Use:   "encode token",
	Short: "Encode JWT token",
	Long: "Encode JWT token, if secret is not provided, default secret is used. " +
		"Pass the token on command line",
	Example: "./jwt encode --secret test '{\"sub\":\"1234567890\",\"name\":\"John Doe\",\"admin\":true}'",
	Run: func(cmd *cobra.Command, args []string) {
		stat, _ := os.Stdin.Stat()
		if stat.Mode()&os.ModeCharDevice == 0 {
			fmt.Println("Piping in")
			reader := bufio.NewReader(os.Stdin)
			text, _ := reader.ReadString('\n')
			fmt.Println(text)
		}
		secret := secret(cmd)

		if len(args) == 0 {
			_, err := fmt.Fprintln(cmd.OutOrStderr(), "Error: No token provided")
			if err != nil {
				return
			}
			return
		}

		fmt.Printf("=== Encoding JWT token with secret === \"%s\" \n", secret)
		encode, err := pkg.Encode(args[0], secret)
		if err != nil {
			_, err := fmt.Fprintln(cmd.OutOrStderr(), err)
			if err != nil {
				return
			}
		}
		_, err = fmt.Fprintln(cmd.OutOrStderr(), encode)
		if err != nil {
			return
		}

		if file != "" {
			fmt.Println("Writing to file")
			err := pkg.ReadFromFile(file)
			if err != nil {

				fmt.Printf("Error %s\n", err.(*fs.PathError).Err.Error())
				fmt.Printf("Operation %s\n", err.(*fs.PathError).Op)
				fmt.Printf("Cannot read file %s\n", err.(*fs.PathError).Path)
			}
			if key == "" {
				pkg.AddToMap(file, encode)
			} else {
				pkg.AddToMap(key, encode)
			}

			println(pkg.ToJSON())
			err = pkg.WriteToFile(file)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	},
}

func secret(cmd *cobra.Command) string {
	strSecret := cmd.Flag("secret").Value.String()
	if strSecret == "" {
		strSecret = pkg.DEFAULT_SECRET
	}
	return strSecret
}

func init() {
	encodeCmd.Flags().StringVarP(&SecretE, "secret", "s", "", "secret key")
	encodeCmd.Flags().StringVarP(&file, "file", "f", "", "file path")
	encodeCmd.Flags().StringVarP(&key, "key", "k", "", "key")
	encodeCmd.Flags().StringVarP(&signingMethod, "signingmethod", "m", "HS256",
		"signing method HS256 | HS384 | HS512")
}
