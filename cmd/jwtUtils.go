package cmd

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"time"
)

type CustomClaims struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	jwt.StandardClaims
}

type CustomMapClaims struct {
	customClaims map[string]string
	jwt.StandardClaims
}

func sampleStandardClaims() jwt.StandardClaims {
	now := time.Now()
	plusYear := now.AddDate(1, 0, 0)
	return jwt.StandardClaims{
		Audience:  "Recipient",
		ExpiresAt: plusYear.UnixMilli(),
		Id:        "1",
		IssuedAt:  now.UnixMilli(),
		Issuer:    "Sample",
		NotBefore: now.UnixMilli(),
		Subject:   "User",
	}
}

func GenerateSimple(claims map[string]string) (string, *jwt.Token) {
	mapClaims := CustomMapClaims{
		customClaims:   claims,
		StandardClaims: sampleStandardClaims(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, mapClaims)
	signingString, err := token.SignedString([]byte("AllYourBase"))
	if err != nil {
		panic(err)
	}

	return signingString, token
}

func GenerateSymmetric(secretKey string, claims map[string]string) (string, *jwt.Token) {
	mapClaims := CustomMapClaims{
		customClaims:   claims,
		StandardClaims: sampleStandardClaims(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, mapClaims)
	signedString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		panic(err)
	}

	return signedString, token
}

func GenerateSigned() string {
	customClaims := customClaims()
	privateKey, _ := privateAndPublicKeyInMemory()
	jwtWithClaims := jwt.NewWithClaims(jwt.SigningMethodRS512, customClaims)
	fromPEM, err := jwt.ParseRSAPrivateKeyFromPEM(privateKey)
	if err != nil {
		panic(err)
	}
	signedString, err := jwtWithClaims.SignedString(fromPEM)

	return signedString
}

func validate(t jwt.Token) {
	method := t.Method
	fmt.Println(method)
}

func customClaims() CustomClaims {
	now := time.Now()
	datePlus3 := now.AddDate(0, 0, 3)

	return CustomClaims{
		"David",
		"Ulicny",
		jwt.StandardClaims{
			Audience:  "Z",
			ExpiresAt: datePlus3.Unix(),
			Id:        "123",
			IssuedAt:  now.Unix(),
			Issuer:    "asd",
			NotBefore: 123,
			Subject:   "asd",
		},
	}
}

func Parse(tokenString string) *jwt.Token {
	parse, err := jwt.Parse(tokenString, nil)
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
	}

	return ""
}

func CustomMapClaimsToString(s CustomMapClaims) string {
	claims := s.customClaims
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
