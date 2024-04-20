package cmd

import (
	"bufio"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/spf13/cobra"
	"io/fs"
	"jwt/pkg"
	"os"
)

var SecretE string

var file string

var key string

type Encode struct {
	secret           string
	file             string
	key              string
	signingMethod    string
	jwtSigningMethod jwt.SigningMethod
}

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

		if len(args) == 0 {
			_, err := fmt.Fprintln(cmd.OutOrStderr(), "Error: No token provided")
			if err != nil {
				return
			}
			return
		}

		e := NewEncode(cmd.Flag("secret").Value.String(),
			cmd.Flag("file").Value.String(),
			cmd.Flag("key").Value.String(),
			cmd.Flag("signingmethod").Value.String(),
		)

		fmt.Printf("=== Encoding JWT token with secret === \"%s\" \n", e.secret)
		encode, err := pkg.EncodeWithMethod(args[0], e.secret, e.jwtSigningMethod)
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

func NewEncode(secret string, file string, key string, signingMethod string) *Encode {
	e := &Encode{
		secret:        secret,
		file:          file,
		key:           key,
		signingMethod: signingMethod,
	}

	if secret == "" {
		e.secret = pkg.DEFAULT_SECRET
	}

	if signingMethod == "" {
		e.signingMethod = "HS256"
	}
	switch signingMethod {
	case "HS256":
		e.jwtSigningMethod = jwt.SigningMethodHS256
	case "HS384":
		e.jwtSigningMethod = jwt.SigningMethodHS384
	case "HS512":
		e.jwtSigningMethod = jwt.SigningMethodHS512
	default:
		e.jwtSigningMethod = jwt.SigningMethodHS256
	}
	return e
}

func init() {
	encodeCmd.Flags().StringVarP(&SecretE, "secret", "s", "", "secret key")
	encodeCmd.Flags().StringVarP(&file, "file", "f", "", "file path")
	encodeCmd.Flags().StringVarP(&key, "key", "k", "", "key")
	encodeCmd.Flags().StringVarP(&signingMethod, "signingmethod", "m", "HS256",
		"signing method HS256 | HS384 | HS512")
}
