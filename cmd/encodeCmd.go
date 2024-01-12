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
	Long: "Encode JWT token, if secret is not provided, default secret is used." +
		"Pass the token on command line",
	Run: func(cmd *cobra.Command, args []string) {
		stat, _ := os.Stdin.Stat()
		if stat.Mode()&os.ModeCharDevice == 0 {
			fmt.Println("Piping in")
			reader := bufio.NewReader(os.Stdin)
			text, _ := reader.ReadString('\n')
			fmt.Println(text)
		}
		secret := cmd.Flag("secret")
		strSecret := cmd.Flag("secret").Value.String()
		if strSecret == "" {
			strSecret = pkg.DEFAULT_SECRET
		}

		if len(args) == 0 {
			fmt.Fprintln(cmd.OutOrStderr(), "Error: No token provided")
			return
		}

		fmt.Printf("=== Encoding JWT token with secret === \"%s\" \n", secret.Value.String())
		encode, err := pkg.Encode(args[0], secret.Value.String())
		if err != nil {
			fmt.Fprintln(cmd.OutOrStderr(), err)
		}
		fmt.Fprintln(cmd.OutOrStderr(), encode)

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
			werr := pkg.WriteToFile(file)
			if werr != nil {
				fmt.Println(werr)
				return
			}
		}
	},
}

func init() {
	encodeCmd.Flags().StringVarP(&SecretE, "secret", "s", "", "secret key")
	encodeCmd.Flags().StringVarP(&file, "file", "f", "", "file path")
	encodeCmd.Flags().StringVarP(&key, "key", "k", "", "key")
}
