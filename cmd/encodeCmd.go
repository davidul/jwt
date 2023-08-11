package cmd

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"jwt/pkg"
	"os"
)

var SecretE string

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
	},
}

func init() {
	encodeCmd.Flags().StringVarP(&SecretE, "secret", "s", "", "secret key")
}
