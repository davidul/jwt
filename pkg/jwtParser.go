package pkg

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt"
	"go.uber.org/zap"
	"strings"
)

type Parser struct {
	tokenString string
	secrets     string
	header      string
	claims      string
	signature   string
	headerMap   map[string]interface{}
	claimsMap   map[string]interface{}
	token       *jwt.Token
}

func NewParser() *Parser {
	logger.Info("Creating new parser")
	return &Parser{}
}

func (p *Parser) SplitToken(tokenString string) ([]string, error) {
	logger.Info("Splitting token")
	if tokenString == "" {
		return nil, fmt.Errorf("token string is empty")
	}
	count := strings.Count(tokenString, ".")
	if count > 2 {
		return nil, fmt.Errorf("token string has more than 2 dots")
	}
	split := strings.Split(tokenString, ".")
	return split, nil
}

func (p *Parser) ParseWithoutVerification(tokenString string) error {
	logger.Info("Parsing token without verification")
	token, err := p.SplitToken(tokenString)
	if err != nil {
		logger.Error("Error splitting token", zap.Error(err))
		return err
	}
	header, claims, signature := token[0], token[1], token[2]
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
	p.tokenString = tokenString
	p.header = string(headerDecoded)
	p.claims = string(claimsDecoded)
	p.signature = string(decodeString)

	err = json.NewDecoder(strings.NewReader(p.header)).Decode(&p.headerMap)
	if err != nil {
		logger.Error("Error decoding header", zap.Error(err))
	}
	err = json.NewDecoder(strings.NewReader(p.claims)).Decode(&p.claimsMap)
	if err != nil {
		logger.Error("Error decoding claims", zap.Error(err))
	}
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
