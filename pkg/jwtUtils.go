package pkg

import (
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt"
	"time"
)

const DEFAULT_SECRET = "AllYourBase"

type CustomMapClaims struct {
	CustomClaims map[string]string
	jwt.StandardClaims
}

// ToMapClaims converts map[string]string to jwt.MapClaims
func ToMapClaims(claims map[string]string) jwt.MapClaims {
	m := make(map[string]interface{})
	for k, v := range claims {
		m[k] = v
	}

	return m
}

// StandardClaimsToMapClaims converts jwt.StandardClaims to jwt.MapClaims
func StandardClaimsToMapClaims(claims jwt.StandardClaims) jwt.MapClaims {
	m := make(map[string]interface{})
	m["aud"] = claims.Audience
	m["exp"] = claims.ExpiresAt
	m["iat"] = claims.IssuedAt
	m["iss"] = claims.Issuer
	m["nbf"] = claims.NotBefore
	m["sub"] = claims.Subject

	return m
}

// sampleStandardClaims returns sample jwt.StandardClaims
// populated with sample values
func sampleStandardClaims() jwt.StandardClaims {
	now := time.Now()
	plusYear := now.AddDate(1, 0, 0)
	minusDay := now.AddDate(0, 0, -1)
	minus2days := now.AddDate(0, 0, -2)
	return jwt.StandardClaims{
		Audience:  "aud",
		ExpiresAt: plusYear.Unix(),
		Id:        "1",
		IssuedAt:  minus2days.Unix(),
		Issuer:    "iss",
		NotBefore: minusDay.Unix(),
		Subject:   "sub",
	}
}

// GenerateSimple alias for GenerateSymmetric
func GenerateSimple(claims map[string]string, signingMethod jwt.SigningMethod) (string, *jwt.Token) {
	return GenerateSymmetric(DEFAULT_SECRET, claims, signingMethod)
}

// GenerateSymmetric generates simple token
// sample claims are populated with sample values and
// optional claims are added to sample claims.
// Default secret is used if none provided.
// returns signed string and token struct
func GenerateSymmetric(secretKey string, claims map[string]string, signingMethod jwt.SigningMethod) (string, *jwt.Token) {
	toMapClaims := StandardClaimsToMapClaims(sampleStandardClaims())
	for k, v := range claims {
		toMapClaims[k] = v
	}

	token := jwt.NewWithClaims(signingMethod, toMapClaims)
	signedString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		panic(err)
	}

	return signedString, token
}

// GenerateSigned generates signed token with private key
func GenerateSigned(claims map[string]string, privateKey *rsa.PrivateKey) string {
	toMapClaims := StandardClaimsToMapClaims(sampleStandardClaims())
	for k, v := range claims {
		toMapClaims[k] = v
	}

	jwtWithClaims := jwt.NewWithClaims(jwt.SigningMethodRS512, toMapClaims)
	signedString, err := jwtWithClaims.SignedString(privateKey)
	if err != nil {
		fmt.Println(err)
	}

	return signedString
}

// Parse parses token string with secret.
// Secret is optional, it is only for validation
func Parse(tokenString string, secret string) *jwt.Token {
	parse, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		fmt.Println(err)
	}

	return parse
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

func ParseWithPublicKey(tokenString string, publicKey *rsa.PublicKey) *jwt.Token {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})
	if err != nil {
		fmt.Println(err)
	}

	return token
}

func HeaderToString(token *jwt.Token) string {
	header := token.Header
	fmt.Println("Header")
	for k, v := range header {
		fmt.Printf("\t%s : %s \n", k, v)
	}

	switch v := token.Claims.(type) {
	case jwt.StandardClaims:
		return StandardClaimsToString(v)
	case CustomMapClaims:
		return CustomMapClaimsToString(v)
	case jwt.MapClaims:
		MapClaimsToString(v)
	}

	return ""
}

func CustomMapClaimsToString(s CustomMapClaims) string {
	claims := s.CustomClaims
	var ret string = "Custom Claims\n"
	for k, v := range claims {
		ret += fmt.Sprintf("\t %s : %s \n", k, v)
	}

	ret += StandardClaimsToString(s.StandardClaims)
	return ret
}

func StandardClaimsToString(s jwt.StandardClaims) string {
	var ret string = "Standard Claims\n"
	ret += fmt.Sprintf("\t Id: %s \n", s.Id)
	ret += fmt.Sprintf("\t Audience: %s\n", s.Audience)
	ret += fmt.Sprintf("\t Issuer: %s\n", s.Issuer)
	ret += fmt.Sprintf("\t Issued at: %s\n", time.UnixMilli(s.IssuedAt).Format(time.RFC3339))
	ret += fmt.Sprintf("\t Not Before: %s\n", time.UnixMilli(s.NotBefore).Format(time.RFC3339))
	ret += fmt.Sprintf("\t Expires at: %s\n", time.UnixMilli(s.ExpiresAt).Format(time.RFC3339))

	return ret
}

func MapClaimsToString(s jwt.MapClaims) string {
	i := s["CustomClaims"]
	if i != nil {
		m := i.(map[string]interface{})
		fmt.Printf("Custom Claims:\n")
		for k, v := range m {
			fmt.Printf("\t%s : %s \n", k, v)
		}
	}

	fmt.Printf("Standard Claims:\n")
	for k, v := range s {
		if k != "CustomClaims" {
			if k == "exp" || k == "iat" || k == "nbf" {
				switch v.(type) {
				case int64:
					milli := time.UnixMilli(v.(int64))
					fmt.Printf("\t%s : %s \n", k, milli.Format(time.RFC3339))
				case float64:
					milli := time.UnixMilli(int64(v.(float64)))
					fmt.Printf("\t%s : %s \n", k, milli.Format(time.RFC3339))
				}

			} else {
				fmt.Printf("\t%s : %s \n", k, v)
			}
		}
	}
	return ""
}

func Encode(data string, secret string) (string, error) {
	c := new(jwt.MapClaims) //map[string]any{}
	err := json.Unmarshal([]byte(data), &c)
	if err != nil {
		return "", err
	}

	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = c
	signingString, err := token.SignedString([]byte(secret))

	return signingString, err
}
