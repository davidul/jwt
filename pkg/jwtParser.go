package pkg

import (
	"encoding/base64"
	"fmt"
	"github.com/golang-jwt/jwt"
	"go.uber.org/zap"
	"strings"
)

func SplitToken(tokenString string) (string, string, string) {
	logger.Info("Splitting token")
	split := strings.Split(tokenString, ".")
	return split[0], split[1], split[2]
}

func ParseWithoutVerification(tokenString string) error {
	logger.Info("Parsing token without verification")
	header, claims, signature := SplitToken(tokenString)
	headerDecoded, err := DecodeBase64String(header)
	if err != nil {
		logger.Error("Error decoding header", zap.Error(err))
		return err
	}

	fmt.Println("Header: ", string(headerDecoded))
	claimsDecoded, err := DecodeBase64String(claims)
	fmt.Println("Claims: ", string(claimsDecoded))
	if err != nil {
		logger.Error("Error decoding claims", zap.Error(err))
		return err
	}
	decodeString, err := DecodeBase64String(signature)
	fmt.Println("Signature: ", decodeString)
	return nil
}

func DecodeBase64String(encoded string) ([]byte, error) {
	logger.Info("Decoding base64 string")
	decoded, err := base64.RawStdEncoding.DecodeString(encoded)
	if err != nil {
		logger.Error("Error decoding base64 string", zap.Error(err))
		return nil, err
	}

	return decoded, nil
}

// Parse parses token string with secret.
// Secret is optional, it is only for validation
func Parse(tokenString string, secret string) (*jwt.Token, error) {
	logger.Info("Parsing token")
	parse, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		logger.Error("Error parsing token", zap.Error(err))
		return nil, err
	}

	return parse, nil
}

func ParseWithPublicKeyFile(tokenString string, publicKeyPath string) *jwt.Token {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		pemBlock := DecodePublicPemFromFile(publicKeyPath)
		publicRsa := UnmarshalPublicRsa(pemBlock)
		return publicRsa, nil
	})
	if err != nil {
		fmt.Println(err)
	}

	return token
}
