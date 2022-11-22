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

func GenerateSimple(claims map[string]string) string {
	mapClaims := CustomMapClaims{
		customClaims:   claims,
		StandardClaims: sampleStandardClaims(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, mapClaims)
	fmt.Println(ToString(token))
	signingString, err := token.SignedString([]byte("AllYourBase"))
	if err != nil {
		panic(err)
	}

	return signingString
}

func GenerateSymmetric(secretKey string, claims map[string]string) string {
	//claims := customClaims()

	withClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, nil)
	signedString, err := withClaims.SignedString([]byte(secretKey))
	if err != nil {
		panic(err)
	}

	return signedString
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
	//duration := time.Duration(30 * time.Second)  //nanosecond
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

func ToString(token *jwt.Token) string {
	header := token.Header
	fmt.Println("Header")
	for k, v := range header {
		fmt.Printf("%s : %s \n", k, v)
	}

	switch v := token.Claims.(type) {
	case jwt.StandardClaims:
		return StandardClaimsToString(v)
	case CustomMapClaims:
		return customMapClaimsToString(v)
	}

	return ""
}

func customMapClaimsToString(s CustomMapClaims) string {
	claims := s.customClaims
	var ret string
	for k, v := range claims {
		ret += fmt.Sprintf("%s : %s \n", k, v)
	}

	return ret
}

func StandardClaimsToString(s jwt.StandardClaims) string {
	b := fmt.Sprintf("Id: %s \nAudience: %s\n", s.Id, s.Audience)
	return b
}
