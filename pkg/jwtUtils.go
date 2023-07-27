package pkg

import (
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt"
	"time"
)

var DEFAULT_SECRET = "AllYourBase"

type CustomMapClaims struct {
	CustomClaims map[string]string
	jwt.StandardClaims
}

func sampleStandardClaims() jwt.StandardClaims {
	now := time.Now()
	plusYear := now.AddDate(1, 0, 0)
	minusDay := now.AddDate(0, 0, -1)
	minus2days := now.AddDate(0, 0, -2)
	return jwt.StandardClaims{
		Audience:  "Recipient",
		ExpiresAt: plusYear.Unix(),
		Id:        "1",
		IssuedAt:  minus2days.Unix(),
		Issuer:    "Sample",
		NotBefore: minusDay.Unix(),
		Subject:   "User",
	}
}

func GenerateSimple(claims map[string]string, signingMethod jwt.SigningMethod) (string, *jwt.Token) {
	mapClaims := CustomMapClaims{
		CustomClaims:   claims,
		StandardClaims: sampleStandardClaims(),
	}

	token := jwt.NewWithClaims(signingMethod, mapClaims)
	signingString, err := token.SignedString([]byte(DEFAULT_SECRET))
	if err != nil {
		panic(err)
	}

	return signingString, token
}

func GenerateSymmetric(secretKey string, claims map[string]string, signingMethod jwt.SigningMethod) (string, *jwt.Token) {
	mapClaims := CustomMapClaims{
		CustomClaims:   claims,
		StandardClaims: sampleStandardClaims(),
	}
	token := jwt.NewWithClaims(signingMethod, mapClaims)
	signedString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		panic(err)
	}

	return signedString, token
}

func GenerateSigned(claims map[string]string) string {
	mapClaims := CustomMapClaims{
		CustomClaims:   claims,
		StandardClaims: sampleStandardClaims(),
	}

	//privateKey, _ := cmd.PrivateAndPublicKeyInMemory()
	private, public := GenKeysRsa()
	fmt.Println("Private and public keys")
	fmt.Println(private)
	fmt.Println(public)
	jwtWithClaims := jwt.NewWithClaims(jwt.SigningMethodRS512, mapClaims)
	//fromPEM, err := jwt.ParseRSAPrivateKeyFromPEM(privateKey)
	//if err != nil {
	//	panic(err)
	//}
	signedString, err := jwtWithClaims.SignedString(private)
	if err != nil {
		fmt.Println(err)
	}

	return signedString
}

func validate(t jwt.Token) {
	method := t.Method
	fmt.Println(method)
}

func Parse(tokenString string, secret string) *jwt.Token {
	parse, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		fmt.Println(err)
	}

	return parse
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
	m := i.(map[string]interface{})
	fmt.Printf("Custom Claims:\n")
	for k, v := range m {
		fmt.Printf("\t%s : %s \n", k, v)
	}

	fmt.Printf("Standard Claims:\n")
	for k, v := range s {
		if k != "CustomClaims" {
			if k == "exp" || k == "iat" || k == "nbf" {
				milli := time.UnixMilli(int64(v.(float64)))

				fmt.Printf("\t%s : %s \n", k, milli.Format(time.RFC3339))
			} else {
				fmt.Printf("\t%s : %s \n", k, v)
			}
		}
	}
	return ""
}

func Encode(data string, secret string) string {
	c := new(jwt.MapClaims) //map[string]any{}
	err := json.Unmarshal([]byte(data), &c)
	if err != nil {
		fmt.Println(err)
	}
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = c
	signingString, err := token.SignedString([]byte(secret))
	if err != nil {
		fmt.Println(err)
	}
	return signingString
}
